package main

import "testing"

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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := parseOptions(tt.args)
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
			if true {
				t.Errorf("parseOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
