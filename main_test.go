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
		{
			name:    "no args",
			args:    []string{},
			want:    &cfg{},
			wantErr: true,
		},
		{
			name:    "correct color option",
			args:    []string{"--color=auto"},
			want:    &cfg{displayCfg: displayCfg{Color: "auto"}},
			wantErr: false,
		},
		{
			name:    "incorrect color option",
			args:    []string{"--color=gibberish"},
			want:    &cfg{},
			wantErr: true,
		},
		{
			name:    "--context overrides both before and after context",
			args:    []string{"-B=1", "-A=4", "--context=3"},
			want:    &cfg{displayCfg: displayCfg{BeforeContext: 3, AfterContext: 3}},
			wantErr: false,
		},
		{
			name:    "--context overrides both before and after context",
			args:    []string{"-B=2", "-A=4"},
			want:    &cfg{displayCfg: displayCfg{BeforeContext: 2, AfterContext: 4}},
			wantErr: false,
		},
		{
			name:    "-C alias behaves the same as --context",
			args:    []string{"-C=5"},
			want:    &cfg{displayCfg: displayCfg{BeforeContext: 5, AfterContext: 5}},
			wantErr: false,
		},
		{
			// positional pattern arg with no flags should not error
			name:    "one positonal arg and no flags",
			args:    []string{"foo"},
			want:    &cfg{},
			wantErr: false,
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
				t.Errorf("Test '%s', parseOptions() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
