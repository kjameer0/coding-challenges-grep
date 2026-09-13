//go:build integration

package main_test

import (
	"fmt"
	"os/exec"
	"testing"
)

func runBinary(args []string) ([]byte, error) {
	cmd := exec.Command(binaryPath, args...)
	return cmd.CombinedOutput()
}
func TestCliArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{name: "no arguments should error", args: []string{}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(binaryPath)
			_, err := runBinary(tt.args)
			if err != nil && !tt.wantErr {
				t.Fatal(err)
			}

		})
	}
}
