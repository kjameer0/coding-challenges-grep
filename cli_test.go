//go:build integration

package main_test

import (
	"os/exec"
	"strings"
	"testing"
)

func TestHelpMessage(t *testing.T) {
	//TODO: determine if this test is worth it
	command := exec.Command("./grep.coding.com")
	b, err := command.Output()
	if err != nil {
		t.Fatal("Command failed to execute")
	}
	if strings.Contains(string(b), "help"){

	}
}
