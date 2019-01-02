package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// VERBOSE turn on logging
var VERBOSE = false

func installHook(binaryPath, installPath string) {
	// check if the install path is a git repository
	gitDirectoryPath := filepath.Join(installPath, ".git")
	_, err := os.Stat(gitDirectoryPath)
	if os.IsNotExist(err) {
		log.Fatal(installPath, "is not a git repository")
	}
	// find the hooks folder
	destinationPath := filepath.Join(gitDirectoryPath, "hooks", "pre-commit")
	destination, _ := os.OpenFile(destinationPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0775) // TODO err
	defer destination.Close()
	// drop the binary into the hooks folder
	binary, err := os.Open(binaryPath) // probably safe
	if err != nil {
		log.Fatal(err)
	}
	defer binary.Close()
	log.Printf("about to copy %s into %s\n", binaryPath, destinationPath)
	nBytes, err := io.Copy(destination, binary)
	if err != nil {
		log.Fatal(err)
	}
	if nBytes == 0 {
		log.Fatal("install failed, no bytes copied")
	}
}

func checkDiff(filename string) {
	var cmd *exec.Cmd
	truffle := regexp.MustCompile(`^\+.*[\#|\/\/]\s?truffle`)
	log.Println("checking " + filename)
	args := []string{"diff", "--unified=0", "--staged", filename}
	cmd = exec.Command("git", args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Run()
	commandOutput := stdout.String()
	splitCommandOutput := strings.Split(commandOutput, "\n")
	for _, line := range splitCommandOutput {
		match := truffle.MatchString(line)
		if match {
			fmt.Println("## TRUFFLE ##")
			fmt.Println("The following line triggered the hook")
			fmt.Println(filename+":", line[1:])
			fmt.Println("## ####### ##")
			os.Exit(56)
		}
	}
}

func checkCommit() {
	var cmd *exec.Cmd
	args := []string{"diff", "--cached", "--diff-filter", "ACMU", "--name-only"}
	cmd = exec.Command("git", args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Run()
	commandOutput := stdout.String()
	splitCommandOutput := strings.Split(commandOutput, "\n")
	for _, filename := range splitCommandOutput {
		if filename == "" {
			continue
		}
		checkDiff(filename)
	}
}

func main() {
	if !VERBOSE {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}
	helpFlag := flag.Bool("h", false, "Display this help text")
	installFlag := flag.Bool("i", false, "Run in install mode, provide install path as positional arg")
	flag.Parse()
	positionalArgs := flag.Args()

	executablePath, _ := os.Executable()
	binaryDir, err := filepath.Abs(filepath.Dir(executablePath))
	if err != nil {
		log.Fatal(err)
	}
	binaryPath := filepath.Join(binaryDir, os.Args[0])

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if *installFlag {
		if flag.NArg() < 1 {
			flag.Usage()
			os.Exit(1)
		}
		installPath := positionalArgs[0]
		log.Println("Installing " + binaryPath + " into " + installPath)
		installHook(binaryPath, installPath)
		log.Println("~Success~")
		os.Exit(0)
	}

	// no flags means we're being called by git
	checkCommit()
	os.Exit(0)
}
