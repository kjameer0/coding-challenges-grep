package stringsearch_test

import (
	"reflect"
	"testing"

	stringsearch "grep.coding.com/stringSearch"
)

func TestBasicRegexSearch_Search(t *testing.T) {
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
		// {
		// 	name:     "single match",
		// 	line:     "hello world",
		// 	patterns: []string{"world"},
		// 	want: []*stringsearch.SearchResult{
		// 		{StartColumn: 6, EndColumn: 11},
		// 	},
		// 	wantErr: false,
		// },
		// {
		// 	name:     "multiple matches of same pattern",
		// 	line:     "foo bar foo baz foo",
		// 	patterns: []string{"foo"},
		// 	want: []*stringsearch.SearchResult{
		// 		{StartColumn: 0, EndColumn: 3},
		// 		{StartColumn: 8, EndColumn: 11},
		// 		{StartColumn: 16, EndColumn: 19},
		// 	},
		// 	wantErr: false,
		// },
		// {
		// 	name:     "no match",
		// 	line:     "hello world",
		// 	patterns: []string{"xyz"},
		// 	want:     []*stringsearch.SearchResult{},
		// 	wantErr:  false,
		// },
		// {
		// 	name:     "multiple patterns",
		// 	line:     "hello world",
		// 	patterns: []string{"hello", "world"},
		// 	want: []*stringsearch.SearchResult{
		// 		{StartColumn: 0, EndColumn: 5},
		// 		{StartColumn: 6, EndColumn: 11},
		// 	},
		// 	wantErr: false,
		// },
		// {
		// 	name:     "repeated chars",
		// 	line:     "aaaaaa",
		// 	patterns: []string{"aaa", "aa"},
		// 	want: []*stringsearch.SearchResult{
		// 		{StartColumn: 0, EndColumn: 3},
		// 		{StartColumn: 3, EndColumn: 6},
		// 		{StartColumn: 0, EndColumn: 2},
		// 		{StartColumn: 2, EndColumn: 4},
		// 		{StartColumn: 4, EndColumn: 6},
		// 	},
		// 	wantErr: false,
		// },
		// {
		// 	// Divergence from FixedStringSearch: "+" is a quantifier here, not a
		// 	// literal, so "a+" matches the leading "a" only rather than the whole
		// 	// two-character line.
		// 	name:     "treats regex metacharacters as operators",
		// 	line:     "a+",
		// 	patterns: []string{"a+"},
		// 	want: []*stringsearch.SearchResult{
		// 		{StartColumn: 0, EndColumn: 1},
		// 	},
		// 	wantErr: false,
		// },
		// {
		// 	// Divergence from FixedStringSearch: "." is any character, so a pattern
		// 	// that would not appear literally still matches.
		// 	name:     "wildcard matches any character",
		// 	line:     "hello world",
		// 	patterns: []string{"h.llo"},
		// 	want: []*stringsearch.SearchResult{
		// 		{StartColumn: 0, EndColumn: 5},
		// 	},
		// 	wantErr: false,
		// },
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
			line:             " Aa ",
			patterns:         []string{"aa"},
			want:             []*stringsearch.SearchResult{
				{StartColumn: 1, EndColumn: 3},
			},
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
			name:             "line regexp anchors the whole pattern",
			line:             "aa ",
			patterns:         []string{"a+"},
			want:             []*stringsearch.SearchResult{},
			extraRegexOption: stringsearch.LineRegexp,
			wantErr:          false,
		},
		{
			name:     "line matching regexp handles regex line syntax",
			line:     "b ",
			patterns: []string{"^b $"},
			want: []*stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 2},
			},
			extraRegexOption: stringsearch.LineRegexp,
			wantErr:          false,
		},
		{
			name:     "partial end matching gets completed by line regex flag",
			line:     "b ",
			patterns: []string{"^b "},
			want: []*stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 2},
			},
			extraRegexOption: stringsearch.LineRegexp,
			wantErr:          false,
		},
		{
			name:     "word matching regexp handles regex line syntax",
			line:     "b ",
			patterns: []string{"^b $"},
			want: []*stringsearch.SearchResult{
				{StartColumn: 0, EndColumn: 2},
			},
			extraRegexOption: stringsearch.WordRegexp,
			wantErr:          false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := stringsearch.NewSearcher(
				stringsearch.WithPatterns(tt.patterns),
				stringsearch.WithSearchType(stringsearch.BasicRegexSearchStrategy),
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
