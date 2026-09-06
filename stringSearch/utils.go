package stringsearch

import (
	"fmt"
)

const IgnoreCaseRegex = "(?i)"

func applySurroundedRegexpChar(pattern string, option ExtraRegexOption, strategy SearchStrategyValue) string {
	switch option {
	case NoExtraRegex:
		return pattern
	case WordRegexp:
		return fmt.Sprintf("\\b%s\\b", pattern)
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
