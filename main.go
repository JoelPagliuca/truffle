package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// VERBOSE turn on logging
var VERBOSE = true

// maybe use `git diff --unified=0 --staged filename`
func checkFile(filename string) {
	log.Println("checking " + filename)
	nocommit := regexp.MustCompile(`.*[\#|\/\/]\s?nocommit`)
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		match := nocommit.MatchString(scanner.Text())
		if match {
			fmt.Println("## NOCOMMIT ##")
			fmt.Println("The following line triggered the hook")
			fmt.Println(filename+":", scanner.Text())
			fmt.Println("## ######## ##")
			os.Exit(1)
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
		checkFile(filename)
	}
}

func main() {
	if !VERBOSE {
		log.SetFlags(0)
	}
	helpFlag := flag.Bool("h", false, "Display this help text")
	installFlag := flag.Bool("i", false, "Run in install mode, provide install path as positional arg")
	flag.Parse()
	positionalArgs := flag.Args()

	binaryPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if *installFlag {
		installPath := positionalArgs[0]
		fmt.Println("Installing " + binaryPath + " into " + installPath)
		// TODO: actually do this
		os.Exit(0)
	}

	// no flags means we're being called by git
	checkCommit()
	os.Exit(0)
}
