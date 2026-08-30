package stringsearch

import "regexp"

func (s *SearchConfig) BasicRegexSearch(line string) ([]*SearchResult, error) {
	results := []*SearchResult{}
	for _, pattern := range s.patterns {
		reg, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		output := reg.FindAllStringIndex(line, -1)
		for _, indexPair := range output {
			results = append(results, NewSearchResult(indexPair[0], indexPair[1]))
		}
	}
	return results, nil
}
