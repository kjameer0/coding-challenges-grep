package stringsearch

import (
	"slices"
	"strings"
)

type fixedStringSearch struct {
	patterns []string
	config   *SearchConfig
}

func (s *fixedStringSearch) applyIgnoreCase(flag bool) {
	if !flag {
		return
	}
	for idx, pattern := range s.patterns {
		s.patterns[idx] = strings.ToLower(pattern)
	}
}

func (s *fixedStringSearch) applyExtraRegexFilters(filter ExtraRegexOption) {
	switch filter {
	case WordRegexp:
		break
	case LineRegexp:
		break
	default:
		break
	}
}

// adjust an existing pattern so that it matches the same expression but with a word break
func addWordRegexpFixed(pattern string) string {
	return ""
}

func NewFixedSearchStrategy(config *SearchConfig) SearchStrategy {
	searcher := &fixedStringSearch{patterns: slices.Clone(config.patterns), config: config}

	searcher.applyExtraRegexFilters(config.ExtraFilter)
	searcher.applyIgnoreCase(config.IgnoreCase)

	return searcher
}

func (s *fixedStringSearch) Search(line string) (*SearchResult, error) {
	if s.config.ExtraFilter == LineRegexp {
		return  line == 
	}
	for _, pattern := range s.patterns {

	}
}
