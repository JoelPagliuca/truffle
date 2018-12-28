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
	"strings"
)

func checkFile(filename string) {
	fmt.Println("checking " + filename)
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
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
		os.Exit(0)
	}

	// no flags means we're being called by git
	checkCommit()
	os.Exit(1)
}
