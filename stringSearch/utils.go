package stringsearch

import "fmt"

const IgnoreCaseRegex = "(?i)"

func applySurroundedRegexpChar(pattern string, option ExtraRegexOption) string {
	switch option {
	case NoExtraRegex:
		return pattern
	case WordRegexp:
		return fmt.Sprintf("\\b%s\\b", pattern)
	case LineRegexp:
		return fmt.Sprintf("^%s$", pattern)
	default:
		return pattern
	}
}
