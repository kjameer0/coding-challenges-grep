package main

import (
	"flag"
	"testing"
)

func Test_parseOptions(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		args    []string
		want    *cfg
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			// no args
			args:    []string{},
			want:    &cfg{},
			wantErr: true,
		},
		{
			// correct color option
			args:    []string{"--color=auto"},
			want:    &cfg{displayCfg: displayCfg{Color: "auto"}},
			wantErr: false,
		},
		{
			// incorrect color option
			args:    []string{"--color=gibberish"},
			want:    &cfg{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := flag.NewFlagSet("testcustomgrep", flag.ExitOnError)
			got, gotErr := parseOptions(tt.args, fs)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("parseOptions() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("parseOptions() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if !isCfgEqual(tt.want, got) {
				t.Errorf("parseOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
