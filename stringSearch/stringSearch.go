package stringsearch

import (
	"regexp"
)

// [][2] represents a slice of size 2 arrays(pair) where pair[0] is the start of the pattern match and pair[1] is the end. (pair[1] + 1) = first idx after the match that is not included in the match. idx in this case refers to bytes, not runes
type LineSearcher interface {
	Search(line string, patterns []string) ([]SearchResult, error)
}

type FixedStringSearch struct {
}

// TODO: in main package make sure there is logic to filter out repeat patterns
func (s *FixedStringSearch) Search(line string, patterns []string) ([]SearchResult, error) {
	results := make([]SearchResult, 0, 100)
	for _, pattern := range patterns {
		reg, err := regexp.Compile(pattern)
		reg.FindAllStringIndex(line)
		if err != nil {
			return nil, err
		}
	}
}

type BasicRegexSearcher struct {
}

/*
I need to be able to receive a string line, patterns, and options and be able to return the actual matches(indices in a string where matches are located).
*/
//is it preferable to have a series of results packed
//the line is assumed because the caller can know what the line is from when they call
type SearchResult struct {
	StartColumn int
	//should not be last idx of match + 1
	EndColumn int
	// TODO: decide whether or not to include the actual string
}

// func isMatch method
// func getMatchText needs to take a searchResult and a line and return the text of the match within the line
func SearchLine(line string, searchStrategy LineSearcher, patterns []string) {

}
