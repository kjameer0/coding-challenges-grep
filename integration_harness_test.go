//go:build integration

package main_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var binaryName = "grep.coding.com"

var binaryPath = ""

func TestMain(m *testing.M) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("could not get current dir: %v", err)
	}

	binaryPath = filepath.Join(dir, binaryName)
	fmt.Println("running program****")
	fmt.Println(binaryPath)
	fmt.Println("running program****")

	m.Run()
}
