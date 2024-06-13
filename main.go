package main

import (
	"flag"
	"os"

	"github.com/Method-Security/codeanalyze/cmd"
)

var version = "none"

func main() {
	flag.Parse()

	codeAnalyze := cmd.NewCodeAnalyze(version)
	codeAnalyze.InitRootCommand()
	codeAnalyze.InitSastCommand()

	if err := codeAnalyze.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
