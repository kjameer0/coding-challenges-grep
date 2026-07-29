package stringsearch

import (
	"regexp"
	"slices"
	"strings"
)

type ExtraRegexOption int

const (
	NoExtraRegex ExtraRegexOption = iota
	WordRegexp
	LineRegexp
)

type SearchConfig struct {
	IgnoreCase  bool
	ExtraFilter ExtraRegexOption
}

type SearchOption func(*SearchConfig)

func WithIgnoreCase(on bool) SearchOption {
	// return a function that can be called with desired option
	return func(s *SearchConfig) { s.IgnoreCase = on }
}

func WithExtraRegexFilter(regexOption ExtraRegexOption) SearchOption {
	return func(s *SearchConfig) { s.ExtraFilter = regexOption }
}

func NewSearchConfig(opts ...SearchOption) *SearchConfig {
	c := &SearchConfig{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// [][2] represents a slice of size 2 arrays(pair) where pair[0] is the start of the pattern match and pair[1] is the end. (pair[1] + 1) = first idx after the match that is not included in the match. idx in this case refers to bytes, not runes
type LineSearcher interface {
	Search(line string, patterns []string) ([]*SearchResult, error)
	BuildPatterns(patterns []string)
}

type FixedStringSearch struct{}

func (s *FixedStringSearch) BuildPatterns(patterns []string, config *SearchConfig) {
	//TODO
	finalizedPatterns := make([]regexp.Regexp, len(patterns))
  patterns = slices.Clone(patterns)
	if config.IgnoreCase {
		//iterate patterns and lower case them
		for idx, pattern := range patterns {
			patterns[idx] = strings.ToLower(pattern)
		}
	}
	//what if it's word regexp.

}
hello  
// TODO: in main package make sure there is logic to filter out repeat patterns
func (s *FixedStringSearch) Search(line string, patterns []string) ([]*SearchResult, error) {
	results := []*SearchResult{}
	for _, pattern := range patterns {
		reg, err := regexp.Compile(regexp.QuoteMeta(pattern))
		output := reg.FindAllStringIndex(line, -1)
		if err != nil {
			return nil, err
		}
		for _, indexPair := range output {
			results = append(results, NewSearchResult(indexPair[0], indexPair[1]))
		}
	}
	return results, nil
}

// if multiple patterns end up highlighting the same parts of the string, consolidate those into the smallest possible window
func ReconcileOverlappingMatches(matches []*SearchResult) []*SearchResult {
	if len(matches) == 0 {
		return []*SearchResult{}
	}
	result := []*SearchResult{}
	intervalStart := matches[0].StartColumn
	intervalEnd := matches[0].EndColumn
	for idx := 1; idx < len(matches); idx++ {
		currentMatch := matches[idx]
		if currentMatch.StartColumn > intervalEnd {
			result = append(result, NewSearchResult(intervalStart, intervalEnd))
			intervalStart = currentMatch.StartColumn
			intervalEnd = currentMatch.EndColumn
		} else if currentMatch.EndColumn > intervalEnd {
			intervalEnd = currentMatch.EndColumn
		}
	}
	result = append(result, NewSearchResult(intervalStart, intervalEnd))
	return result
}

type BasicRegexSearch struct{}

func (s *BasicRegexSearch) Search(line string, patterns []string) ([]*SearchResult, error) {
	results := make([]*SearchResult, 0, 100)

	for _, pattern := range patterns {
		r, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		matches := r.FindAllIndex([]byte(line), -1)
		for _, match := range matches {
			results = append(results, NewSearchResult(match[0], match[1]))
		}
	}
	return results, nil
}

type SearchResult struct {
	StartColumn int
	EndColumn   int
	// TODO: decide whether or not to include the actual string
}

func NewSearchResult(startColumn, endColumn int) *SearchResult {
	return &SearchResult{StartColumn: startColumn, EndColumn: endColumn}
}

// func getMatchText needs to take a searchResult and a line and return the text of the match within the line
func SearchLine(line string, searchStrategy LineSearcher, patterns []string, searchConfig *SearchConfig) ([]*SearchResult, error) {
	ignoreCase := searchConfig.IgnoreCase
	extraRegex := searchConfig.ExtraFilter
	results, err := searchStrategy.Search(line, patterns)
	if err != nil {
		return nil, err
	}
	return ReconcileOverlappingMatches(results), nil
}
