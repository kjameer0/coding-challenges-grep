package stringsearch

import (
	"regexp"
)

type FixedStringSearch struct {
	config                  *SearchConfig
	postCompilationPatterns []*regexp.Regexp
}

func NewFixedStringSearch(c *SearchConfig) *FixedStringSearch {
	var fss *FixedStringSearch = &FixedStringSearch{}
	fss.postCompilationPatterns = make([]*regexp.Regexp, 0)
	for _, pattern := range c.patterns {
		pattern = applySurroundedRegexpChar(regexp.QuoteMeta(pattern), c.ExtraFilter)
		if c.IgnoreCase {
			pattern = IgnoreCaseRegex + pattern
		}
		re := regexp.MustCompile(pattern)
		fss.postCompilationPatterns = append(fss.postCompilationPatterns, re)
	}
	return fss
}

func (s *FixedStringSearch) Search(line string) ([]*SearchResult, error) {
	results := []*SearchResult{}
	for _, re := range s.postCompilationPatterns {
		output := re.FindAllStringIndex(line, -1)
		for _, indexPair := range output {
			results = append(results, NewSearchResult(indexPair[0], indexPair[1]))
		}
	}
	return results, nil
}
