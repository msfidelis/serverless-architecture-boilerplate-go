//go:build mage
// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = PHONY

func PHONY() {
	mg.Deps(Clean, Build, Deploy, Remove)
}

// clean remove all bin
func Clean() error {
	return sh.Rm("bin")
}

// build build each module into "bin"
func Build() error {
	mg.Deps(Clean)
	os.Setenv("GO111MODULE", "on")
	os.Setenv("GOOS", "linux")
	gobuild := sh.RunCmd("go", "build", "-ldflags", "-s -w", "-o")

	bins := []struct {
		output string
		source string
	}{
		{"bin/books/endpoints/create", "modules/books/endpoints/create.go"},
		{"bin/books/endpoints/read", "modules/books/endpoints/read.go"},
		{"bin/books/endpoints/detail", "modules/books/endpoints/detail.go"},
		{"bin/books/endpoints/update", "modules/books/endpoints/update.go"},
		{"bin/books/endpoints/delete", "modules/books/endpoints/delete.go"},

		{"bin/books/functions/worker", "modules/books/functions/worker.go"},
	}
	for _, bin := range bins {
		err := gobuild(bin.output, bin.source)
		if err != nil {
			return err
		}
	}
	return nil
}

// test run all go tests
func Test() error {
	return sh.Run("go", "test", "./...")
}

// deploy force deployment to default
func Deploy() error {
	mg.Deps(Clean, Build)
	return sh.Run("serverless", "deploy", "--verbose", "--force")
}

func Remove() error {
	mg.Deps(Clean)
	return sh.Run("serverless", "remove")
}
