package stringsearch

import (
	"fmt"
	"testing"
)

func Test_applySurroundedRegexpChar(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		pattern  string
		option   ExtraRegexOption
		strategy SearchStrategyValue
		want     string
	}{
		{
			pattern:  "^b",
			option:   LineRegexp,
			strategy: BasicRegexSearchStrategy,
			want:     "^b$",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := applySurroundedRegexpChar(tt.pattern, tt.option, tt.strategy)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				fmt.Printf("got length: %d, want length: %d\n", len(got), len(tt.want))
				t.Errorf("applySurroundedRegexpChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
