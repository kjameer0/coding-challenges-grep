package stringsearch_test

import (
<<<<<<< HEAD
=======
	"reflect"
>>>>>>> pattern-match
	"testing"

	stringsearch "grep.coding.com/stringSearch"
)

func TestFixedStringSearch_Search(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		line     string
		patterns []string
<<<<<<< HEAD
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
=======
		want     []stringsearch.SearchResult
		wantErr  bool
	}{
		{
			name:     "single match",
			line:     "hello world",
			patterns: []string{"world"},
			want: []stringsearch.SearchResult{
				{StartColumn: 6, EndColumn: 11},
>>>>>>> pattern-match
			},
			wantErr: false,
		},
		{
<<<<<<< HEAD
			name:     "one pattern, two matches",
			patterns: []string{"hello"},
			line:     "hellohello",
			want: []*stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 5},
				{StartColumn: 5, EndColumn: 10},
=======
			name:     "multiple matches of same pattern",
			line:     "foo bar foo baz foo",
			patterns: []string{"foo"},
			want: []stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 3},
				{StartColumn: 8, EndColumn: 11},
				{StartColumn: 16, EndColumn: 19},
>>>>>>> pattern-match
			},
			wantErr: false,
		},
		{
<<<<<<< HEAD
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
			name:     "multiple patterns",
			patterns: []string{"hh", "h"},
			line:     "hhhh",
			want: []*stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 2},
				{StartColumn: 2, EndColumn: 4},
				{StartColumn: 0, EndColumn: 1},
				{StartColumn: 1, EndColumn: 2},
				{StartColumn: 2, EndColumn: 3},
				{StartColumn: 3, EndColumn: 4},
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
=======
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
		{
			name:     "matches literal text of a regular expression string",
			line:     "a+",
			patterns: []string{"aa"},
			want:     []stringsearch.SearchResult{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
>>>>>>> pattern-match
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
<<<<<<< HEAD
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
					t.Errorf("Name: %s Match %+v not found in wanted results %+v", tt.name, result, wanted)
				}
			}
		})
	}
}

func TestBasicRegexSearch_Search(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		line     string
		patterns []string
		want     []*stringsearch.SearchResult
		wantErr  bool
	}{
		// TODO: Add test cases.
		{
			name:     "regex string pattern searches text for regex pattern",
			patterns: []string{"[a-z]+"},
			line:     "[a-z]+?",
			want: []*stringsearch.SearchResult{
				{
					StartColumn: 1,
					EndColumn:   2,
				},
				{
					StartColumn: 3,
					EndColumn:   4,
				},
			},
			wantErr: false,
		},
		{
			name:     "multiple patterns yield results accordingl",
			patterns: []string{"[a-z]+"},
			line:     "[a-z]+?",
			want: []*stringsearch.SearchResult{
				{
					StartColumn: 1,
					EndColumn:   2,
				},
				{
					StartColumn: 3,
					EndColumn:   4,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s stringsearch.BasicRegexSearch
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

func TestReconcileOverlappingMatches(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		matches []*stringsearch.SearchResult
		want    []*stringsearch.SearchResult
	}{
		// TODO: Add test cases.
		{
			name:    "Empty input yields empty array of match intervals",
			matches: []*stringsearch.SearchResult{},
			want:    []*stringsearch.SearchResult{},
		},
		{
			name: "Overlapping intervals get merged",
			matches: []*stringsearch.SearchResult{
				{
					StartColumn: 0,
					EndColumn:   1,
				},
				{
					StartColumn: 1,
					EndColumn:   2,
				},
			},
			want: []*stringsearch.SearchResult{
				{
					StartColumn: 0,
					EndColumn:   2,
				},
			},
		},
		{
			name: "Multiple non-consecutive overlapping intervals get merged",
			matches: []*stringsearch.SearchResult{
				{
					StartColumn: 0,
					EndColumn:   1,
				},
				{
					StartColumn: 1,
					EndColumn:   2,
				},
				{
					StartColumn: 10,
					EndColumn:   30,
				},
				{
					StartColumn: 25,
					EndColumn:   45,
				},
			},
			want: []*stringsearch.SearchResult{
				{
					StartColumn: 0,
					EndColumn:   2,
				},
				{
					StartColumn: 10,
					EndColumn:   45,
				},
			},
		},
		{
			name: "Intervals inside of a larger interval get subsumed",
			matches: []*stringsearch.SearchResult{
				{
					StartColumn: 0,
					EndColumn:   1,
				},
				{
					StartColumn: 1,
					EndColumn:   2,
				},
				{
					StartColumn: 10,
					EndColumn:   30,
				},
				{
					StartColumn: 15,
					EndColumn:   20,
				},
				{
					StartColumn: 25,
					EndColumn:   45,
				},
			},
			want: []*stringsearch.SearchResult{
				{
					StartColumn: 0,
					EndColumn:   2,
				},
				{
					StartColumn: 10,
					EndColumn:   45,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stringsearch.ReconcileOverlappingMatches(tt.matches)
			// TODO: update the condition below to compare got with tt.want.
			if len(got) != len(tt.want) {
				t.Errorf("Unequal got and want lengths. got %v+, want %v+", got, tt.want)
				return
			}
			for intervalIdx := 0; intervalIdx < len(got); intervalIdx++ {
				gotInterval := got[intervalIdx]
				wantInterval := tt.want[intervalIdx]
				if gotInterval.StartColumn != wantInterval.StartColumn || gotInterval.EndColumn != wantInterval.EndColumn {
					t.Errorf("Name: %s got %v+, want %v+", tt.name, got, tt.want)
				}
=======
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
>>>>>>> pattern-match
			}
		})
	}
}
