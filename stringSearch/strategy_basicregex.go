package stringsearch

import (
	"regexp"
)

type BasicRegexSearch struct {
	config                  *SearchConfig
	postCompilationPatterns []*regexp.Regexp
}

func NewBasicRegexSearch(c *SearchConfig) *BasicRegexSearch {
	var fss *BasicRegexSearch = &BasicRegexSearch{}
	fss.postCompilationPatterns = make([]*regexp.Regexp, 0)
	for _, pattern := range c.patterns {
		pattern = applySurroundedRegexpChar(pattern, c.ExtraFilter, c.SearchType)
		if c.IgnoreCase {
			pattern = IgnoreCaseRegex + pattern
		}
		re := regexp.MustCompile(pattern)
		fss.postCompilationPatterns = append(fss.postCompilationPatterns, re)
	}
	return fss
}

func (s *BasicRegexSearch) Search(line string) ([]*SearchResult, error) {
	results := []*SearchResult{}

	for _, re := range s.postCompilationPatterns {
		output := re.FindAllStringIndex(line, -1)
		for _, indexPair := range output {
			results = append(results, NewSearchResult(indexPair[0], indexPair[1]))
		}
	}
	return results, nil
}
