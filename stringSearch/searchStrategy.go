package stringsearch

type SearchStrategy interface {
	applyIgnoreCase(flag bool)
	applyExtraRegexFilters(filter ExtraRegexOption)
	Search(line string) (*SearchResult, error)
}
