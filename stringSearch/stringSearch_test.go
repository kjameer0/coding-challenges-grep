package stringsearch_test

import (
	"reflect"
	"testing"

	stringsearch "grep.coding.com/stringSearch"
)

func TestFixedStringSearch_Search(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		line     string
		patterns []string
		want     []stringsearch.SearchResult
		wantErr  bool
	}{
		{
			name:     "single match",
			line:     "hello world",
			patterns: []string{"world"},
			want: []stringsearch.SearchResult{
				{StartColumn: 6, EndColumn: 11},
			},
			wantErr: false,
		},
		{
			name:     "multiple matches of same pattern",
			line:     "foo bar foo baz foo",
			patterns: []string{"foo"},
			want: []stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 3},
				{StartColumn: 8, EndColumn: 11},
				{StartColumn: 16, EndColumn: 19},
			},
			wantErr: false,
		},
		{
			name:     "no match",
			line:     "hello world",
			patterns: []string{"xyz"},
			want:     []stringsearch.SearchResult{},
			wantErr:  false,
		},
		{
			name:     "multiple patterns",
			line:     "hello world",
			patterns: []string{"hello", "world"},
			want: []stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 5},
				{StartColumn: 6, EndColumn: 11},
			},
			wantErr: false,
		},
		{
			name:     "repeated chars",
			line:     "aaaaaa",
			patterns: []string{"aaa", "aa"},
			want: []stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 3},
				{StartColumn: 3, EndColumn: 6},
				{StartColumn: 0, EndColumn: 2},
				{StartColumn: 2, EndColumn: 4},
				{StartColumn: 4, EndColumn: 6},
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
