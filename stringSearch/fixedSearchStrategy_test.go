package stringsearch

import "testing"

func Test_fixedStringSearch_Search(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		line       string
		want       *SearchResult
		wantErr    bool
		ignoreCase bool
		extraRegex ExtraRegexOption
		patterns []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			c := &SearchConfig{
				IgnoreCase:  tt.ignoreCase,
				ExtraFilter: tt.extraRegex,
				SearchType:  FixedStringSearchStrategy,
				patterns:    tt.patterns,
			}
			var s SearchStrategy = NewFixedSearchStrategy(c)
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
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
