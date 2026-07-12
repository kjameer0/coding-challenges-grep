package stringsearch_test

import (
	"testing"

	stringsearch "grep.coding.com/stringSearch"
)

func TestFixedStringSearch_Search(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		line     string
		patterns []string
		want     []*stringsearch.SearchResult
		wantErr  bool
	}{
		// TODO: Add test cases)
		{
			name:     "one pattern",
			patterns: []string{"hello"},
			line:     "hello",
			want: []*stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 5},
			},
			wantErr: false,
		},
		{
			name:     "one pattern, two matches",
			patterns: []string{"hello"},
			line:     "hellohello",
			want: []*stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 5},
				{StartColumn: 5, EndColumn: 10},
			},
			wantErr: false,
		},
		{
			name:     "one pattern, does not try for overlapping matches",
			patterns: []string{"hh"},
			line:     "hhhh",
			want: []*stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 2},
				{StartColumn: 2, EndColumn: 4},
			},
			wantErr: false,
		},
		{
			name:     "no pattern",
			patterns: []string{},
			line:     "hhhh",
			want:     []*stringsearch.SearchResult{},
			wantErr:  false,
		},
		{
			name:     "regex string pattern searches for literal expression text",
			patterns: []string{"[a-z]+"},
			line:     "[a-z]+?",
			want: []*stringsearch.SearchResult{
				{
					StartColumn: 0,
					EndColumn:   6,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s stringsearch.FixedStringSearch
			got, gotErr := s.Search(tt.line, tt.patterns)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Search() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Search() succeeded unexpectedly")
			}
			for _, result := range got {
				foundMatch := false
				for _, wantedResult := range tt.want {
					if wantedResult.StartColumn == result.StartColumn && wantedResult.EndColumn == result.EndColumn {
						foundMatch = true
					}
				}
				if !foundMatch {
					wanted := make([]stringsearch.SearchResult, len(tt.want))
					for i, w := range tt.want {
						wanted[i] = *w
					}
					t.Errorf("Match %+v not found in wanted results %+v", result, wanted)
				}
			}
		})
	}
}
