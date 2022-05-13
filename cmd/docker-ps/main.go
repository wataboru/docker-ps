package main

import (
	"flag"
	"fmt"
	"github.com/wataboru/git-fuzzy-find-commit-message/dockerps"
	"os"
)

const (
	// ExitCodeSuccess indicates a successful exit
	ExitCodeSuccess = 0
	// ExitCodeError indicates that some error occurred during processing
	ExitCodeError = 1
	// Version is application version
	Version = "1.0.0"
)

var (
	showVersion bool
)

func init() {
	flag.BoolVar(&showVersion, "v", false, "show version (short)")
	flag.BoolVar(&showVersion, "version", false, "show version")
}

func run() int {
	flag.Parse()
	if showVersion {
		fmt.Println("docker-ps version " + Version)
		return ExitCodeSuccess
	}

	if err := dockerps.Copy(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return ExitCodeError
	}

	return ExitCodeSuccess
}

func main() {
	os.Exit(run())
}
