package fileparse_test

import (
	"path/filepath"
	"reflect"
	"testing"

	fileparse "grep.coding.com/fileParse"
)

func TestParseCfg_ConstructFileSet(t *testing.T) {
	testDirName := "testRootDir"
	simpleRecursionDir := filepath.Join(testDirName, "simpleRecursionRoot")

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		opts    []fileparse.ParseOption
		want    []string
		dirPath string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Files can be discovered without recursing into subdirs",
			want: []string{"f1"},
			opts: []fileparse.ParseOption{
				fileparse.WithUniversePaths([]string{testDirName}),
			},
		},
		{
			want: []string{"child1/f4", "f3"},
			opts: []fileparse.ParseOption{
				fileparse.WithUniversePaths([]string{simpleRecursionDir}),
				fileparse.WithRecursiveResolution(true),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tt.opts = append(tt.opts, fileparse.WithUniversePaths([]string{root}))
			p, err := fileparse.NewParseCfg(tt.opts...)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}

			got, gotErr := p.ConstructFileSet()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ConstructFileSet() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ConstructFileSet() succeeded unexpectedly")
			}

			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("ConstructFileSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
