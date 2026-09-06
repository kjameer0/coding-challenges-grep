package stringsearch_test

import (
	"reflect"
	"testing"

	stringsearch "grep.coding.com/stringSearch"
)

func TestSearchConfig_FixedStringSearch(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		line             string
		patterns         []string
		want             []*stringsearch.SearchResult
		ignoreCase       bool
		extraRegexOption stringsearch.ExtraRegexOption
		wantErr          bool
	}{
		{
			name:     "single match",
			line:     "hello world",
			patterns: []string{"world"},
			want: []*stringsearch.SearchResult{
				{StartColumn: 6, EndColumn: 11},
			},
			wantErr: false,
		},
		{
			name:     "multiple matches of same pattern",
			line:     "foo bar foo baz foo",
			patterns: []string{"foo"},
			want: []*stringsearch.SearchResult{
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
			want:     []*stringsearch.SearchResult{},
			wantErr:  false,
		},
		{
			name:     "multiple patterns",
			line:     "hello world",
			patterns: []string{"hello", "world"},
			want: []*stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 5},
				{StartColumn: 6, EndColumn: 11},
			},
			wantErr: false,
		},
		{
			name:     "repeated chars",
			line:     "aaaaaa",
			patterns: []string{"aaa", "aa"},
			want: []*stringsearch.SearchResult{
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
			want:     []*stringsearch.SearchResult{},
			wantErr:  false,
		},
		{
			// Divergence from https://pkg.go.dev/testing#hdr-MainRegexSearch: "+" is literal here, so the whole
			// two-character line matches.
			name:     "treats regex metacharacters as literals",
			line:     "a+",
			patterns: []string{"a+"},
			want: []*stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 2},
			},
			wantErr: false,
		},
		{
			name:     "matches different cased text when ignoreCase is true",
			line:     "A",
			patterns: []string{"a"},
			want: []*stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 1},
			},
			ignoreCase: true,
			wantErr:    false,
		},
		{
			name:             "matches different cased text when ignoreCase is true, and matches word break regex",
			line:             "Ab",
			patterns:         []string{"a"},
			want:             []*stringsearch.SearchResult{},
			ignoreCase:       true,
			extraRegexOption: stringsearch.WordRegexp,
			wantErr:          false,
		},
		{
			name:             "word regexp matches a standalone word",
			line:             "a b",
			patterns:         []string{"a"},
			want:             []*stringsearch.SearchResult{{StartColumn: 0, EndColumn: 1}},
			extraRegexOption: stringsearch.WordRegexp,
			wantErr:          false,
		},
		{
			name:             "line matching regexp matches the whole line",
			line:             "b ",
			patterns:         []string{"b"},
			want:             []*stringsearch.SearchResult{},
			extraRegexOption: stringsearch.LineRegexp,
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := stringsearch.NewSearcher(
				stringsearch.WithPatterns(tt.patterns),
				stringsearch.WithSearchType(stringsearch.FixedStringSearchStrategy),
				stringsearch.WithIgnoreCase(tt.ignoreCase),
				stringsearch.WithExtraRegexFilter(tt.extraRegexOption),
			)
			if err != nil {
				t.Fatalf("NewSearchConfig() failed: %v", err)
			}
			got, gotErr := s.Search(tt.line)
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

func TestReconcileOverlappingMatches(t *testing.T) {
	tests := []struct {
		name    string
		matches []*stringsearch.SearchResult
		want    []*stringsearch.SearchResult
	}{
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
			if len(got) != len(tt.want) {
				t.Errorf("Unequal got and want lengths. got %v, want %v", got, tt.want)
				return
			}
			for intervalIdx := 0; intervalIdx < len(got); intervalIdx++ {
				gotInterval := got[intervalIdx]
				wantInterval := tt.want[intervalIdx]
				if gotInterval.StartColumn != wantInterval.StartColumn || gotInterval.EndColumn != wantInterval.EndColumn {
					t.Errorf("Name: %s got %v, want %v", tt.name, got, tt.want)
				}
			}
		})
	}
}
