package stringsearch

import "strings"


//[][2] represents a slice of size 2 arrays(pair) where pair[0] is the start of the pattern match and pair[1] is the end. (pair[1] + 1) = first idx after the match that is not included in the match. idx in this case refers to bytes, not runes
type LineSearcher interface {
	Search (line string, patterns []string) [][2]int
}

type SearchOptions struct {
	IgnoreCase bool
	UseExclude bool
}
/*
I need to be able to receive a string line, patterns, and options and be able to return the actual matches(indices in a string where matches are located).
*/
type SearchStrategy struct {

}

//use extended regex
//use fixed string matching
//use Basic regex


