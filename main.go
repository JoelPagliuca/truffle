package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func usage() {
	flag.Usage()
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
		usage()
		os.Exit(0)
	}

	if *installFlag {
		installPath := positionalArgs[0]
		fmt.Println("Installing " + binaryPath + " into " + installPath)
		os.Exit(0)
	}

	// no flags means we're being called by git
	fmt.Println("Being called as a hook, args:", positionalArgs)
	os.Exit(0)
}
