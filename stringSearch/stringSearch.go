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
	results := []SearchResult{}
	for _, pattern := range patterns {
		reg, err := regexp.Compile(regexp.QuoteMeta(pattern))
		output := reg.FindAllStringIndex(line, -1)
		if err != nil {
			return nil, err
		}
		for _, indexPair := range output {
			results = append(results, SearchResult{StartColumn: indexPair[0], EndColumn: indexPair[1]})
		}
	}
	return results, nil
}

type BasicRegexSearcher struct {
}

func (s *BasicRegexSearcher) Search(line string, patterns []string) ([]SearchResult, error) {
	results := []SearchResult{}
	for _, pattern := range patterns {
		reg, err := regexp.Compile(pattern)
		output := reg.FindAllStringIndex(line, -1)
		if err != nil {
			return nil, err
		}
		for _, indexPair := range output {
			results = append(results, SearchResult{StartColumn: indexPair[0], EndColumn: indexPair[1]})
		}
	}
	return results, nil
}


type SearchResult struct {
	StartColumn int
	//should not be last idx of match + 1
	EndColumn int
	// TODO: decide whether or not to include the actual string
}

// func getMatchText needs to take a searchResult and a line and return the text of the match within the line
func SearchLine(line string, searchStrategy LineSearcher, patterns []string) ([]SearchResult, error) {
	return searchStrategy.Search(line,patterns)
}
