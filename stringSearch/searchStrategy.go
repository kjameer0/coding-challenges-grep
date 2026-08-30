package stringsearch

type Searcher interface {
	Search(line string) ([]*SearchResult, error)
}
