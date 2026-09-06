package stringsearch

import (
	"errors"
	"fmt"
)

type ExtraRegexOption int

const (
	NoExtraRegex ExtraRegexOption = iota
	WordRegexp
	LineRegexp
)

type SearchStrategyValue int

const (
	FixedStringSearchStrategy SearchStrategyValue = iota
	BasicRegexSearchStrategy
)

type SearchConfig struct {
	IgnoreCase  bool
	ExtraFilter ExtraRegexOption
	SearchType  SearchStrategyValue
	patterns    []string
}

type SearchOption func(*SearchConfig)

func WithIgnoreCase(on bool) SearchOption {
	// return a function that can be called with desired option
	return func(s *SearchConfig) { s.IgnoreCase = on }
}

func WithExtraRegexFilter(regexOption ExtraRegexOption) SearchOption {
	return func(s *SearchConfig) { s.ExtraFilter = regexOption }
}

func WithSearchType(searchStrategy SearchStrategyValue) SearchOption {
	return func(s *SearchConfig) { s.SearchType = searchStrategy }
}

func WithPatterns(patterns []string) SearchOption {
	return func(s *SearchConfig) { s.patterns = patterns }
}

var InvalidSearchStrategyError = errors.New("Invalid search strategy provided")

func NewSearcher(opts ...SearchOption) (Searcher, error) {
	c := &SearchConfig{
		IgnoreCase:  false,
		ExtraFilter: NoExtraRegex,
		SearchType:  BasicRegexSearchStrategy,
		patterns:    []string{},
	}
	for _, opt := range opts {
		opt(c)
	}
	// new searcher implementations need to be added here to be available to users
	switch c.SearchType {
	case BasicRegexSearchStrategy:
		return NewBasicRegexSearch(c), nil
	case FixedStringSearchStrategy:
		return NewFixedStringSearch(c), nil
	default:
		return nil, InvalidSearchStrategyError
	}
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

type SearchResult struct {
	StartColumn int
	EndColumn   int
	// TODO: decide whether or not to include the actual string
}

func (s *SearchResult) String() string {
	return fmt.Sprintf("Start Column: %d, End Column: %d", s.StartColumn, s.EndColumn)
}
func NewSearchResult(startColumn, endColumn int) *SearchResult {
	return &SearchResult{StartColumn: startColumn, EndColumn: endColumn}
}
