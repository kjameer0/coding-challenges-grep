package optionparse

import (
	"flag"
	"testing"
)

func Test_parseOptions(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		args    []string
		want    *Cfg
		wantErr bool
	}{
		{
			name:    "no args",
			args:    []string{},
			want:    &Cfg{},
			wantErr: true,
		},
		{
			name:    "correct color option",
			args:    []string{"--color=auto", "hello"},
			want:    &Cfg{DisplayCfg: DisplayCfg{Color: "auto"}},
			wantErr: false,
		},
		{
			name:    "correct color option, with alias",
			args:    []string{"--colour=auto", "hello"},
			want:    &Cfg{DisplayCfg: DisplayCfg{Color: "auto"}},
			wantErr: false,
		},
		{
			name:    "incorrect color option",
			args:    []string{"--color=gibberish", "hello"},
			want:    &Cfg{},
			wantErr: true,
		},
		{
			name:    "--context overrides both before and after context",
			args:    []string{"-B=1", "-A=4", "--context=3", "hello"},
			want:    &Cfg{DisplayCfg: DisplayCfg{BeforeContext: 3, AfterContext: 3}},
			wantErr: false,
		},
		{
			name:    "--context overrides both before and after context",
			args:    []string{"-B=2", "-A=4", "pattern"},
			want:    &Cfg{DisplayCfg: DisplayCfg{BeforeContext: 2, AfterContext: 4}},
			wantErr: false,
		},
		{
			name:    "-C alias behaves the same as --context",
			args:    []string{"-C=5", "hello"},
			want:    &Cfg{DisplayCfg: DisplayCfg{BeforeContext: 5, AfterContext: 5}},
			wantErr: false,
		},
		{
			name:    "one positonal arg and no flags",
			args:    []string{"foo"},
			want:    &Cfg{},
			wantErr: false,
		},
		{
			name:    "just non-pattern flags and no postional args",
			args:    []string{"-c"},
			want:    &Cfg{},
			wantErr: true,
		},
		{
			name:    "allows single regexp arg with no positional arg",
			args:    []string{"-e=hello"},
			want:    &Cfg{PatternCfg: PatternCfg{patterns: []string{"hello"}}},
			wantErr: false,
		},
		{
			name:    "allows multiple regexp args with positional arg",
			args:    []string{"-e=hello", "-e", "welcome", "/try/"},
			want:    &Cfg{PatternCfg: PatternCfg{patterns: []string{"hello", "welcome", "/try/"}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := flag.NewFlagSet("testcustomgrep", flag.ExitOnError)
			got, gotErr := ParseOptions(tt.args, fs)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("parseOptions() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("parseOptions() succeeded unexpectedly")
			}

			if !isCfgEqual(tt.want, got) {
				t.Errorf("Test '%s', parseOptions() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
