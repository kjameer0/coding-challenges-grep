package stringsearch

import (
	"fmt"
	"strings"
)

const IGNORE_CASE_REGEX = "(?i)"
const WORD_BREAK_SEQUENCE = "\\b"
func applySurroundedRegexpChar(pattern string, option ExtraRegexOption, strategy SearchStrategyValue) string {
	switch option {
	case NoExtraRegex:
		return pattern
	case WordRegexp:
		if strategy != BasicRegexSearchStrategy {
			return fmt.Sprintf("\\b%s\\b", pattern)
		}
		var leftSide string
		var rightSide string
		if strings.HasPrefix(pattern, "\\b") {
			leftSide = "\\b"
		}
		if strings.HasSuffix(pattern, "\\b") {
			rightSide = "\\b"
		}
		return fmt.Sprintf("%s%s%s", leftSide, pattern, rightSide)
	case LineRegexp:
		if strategy != BasicRegexSearchStrategy {
			return fmt.Sprintf("^%s$", pattern)
		}
		var leftSide string = ""
		var rightSide string = ""
		if pattern[0] != '^' {
			leftSide = "^"
		}
		if len([]rune(pattern))-1 != '$' {
			rightSide = "$"
		}
		return fmt.Sprintf("%s%s%s", leftSide, pattern, rightSide)
	default:
		return pattern
	}
}
