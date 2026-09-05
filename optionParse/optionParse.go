package optionparse

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"

	"golang.org/x/term"
)

type ColorOption string

const (
	COLOR_NEVER       = "never"
	COLOR_ALWAYS      = "always"
	COLOR_AUTO        = "auto"
	COLOR             = "color"
	COLOUR            = "colour"
	CONTEXT           = "context"
	AFTER_CONTEXT     = "after-context"
	BEFORE_CONTEXT    = "before-context"
	COUNT             = "count"
	COUNT_ALIAS       = "c"
	BEFORE_ALIAS      = "B"
	AFTER_ALIAS       = "A"
	CONTEXT_ALIAS     = "C"
	REGEXP            = "regexp"
	REGEXP_ALIAS      = "e"
	WORD_REGEXP       = "word-regexp"
	WORD_REGEXP_ALIAS = "w"
	LINE_REGEXP       = "line-regexp"
	LINE_REGEXP_ALIAS = "x"
	IGNORE_CASE       = "ignore-case"
	IGNORE_CASE_ALIAS = "i"
)

var ContextAliases []string = []string{CONTEXT, CONTEXT_ALIAS}
var CfterContextAliases []string = []string{AFTER_CONTEXT, AFTER_ALIAS}
var CeforeContextAliases []string = []string{BEFORE_CONTEXT, BEFORE_ALIAS}
var ColorAliases []string = []string{COLOR, COLOUR}

func IsColorOption(option string) bool {
	allowedOptions := []string{COLOR_AUTO, COLOR_ALWAYS, COLOR_NEVER}
	return slices.Contains(allowedOptions, option)
}

type Cfg struct {
	IsTerm bool
	FileCfg
	DisplayCfg
	PatternCfg
}

type FileCfg struct {
	//TODO: plan out types for file inclusion ond exclusion
}

type DisplayCfg struct {
	Color         ColorOption
	BeforeContext int
	AfterContext  int
	UseByteOffset bool
	UseCount      bool
}

type PatternCfg struct {
	WordRegexp bool
	LineRegexp bool
	IgnoreCase bool
	patterns   []string
}

func isPatternCfgEqual(patternCfg1 *PatternCfg, patternCfg2 *PatternCfg) bool {
	if patternCfg1.WordRegexp != patternCfg2.WordRegexp {
		return false
	}
	if patternCfg1.LineRegexp != patternCfg2.LineRegexp {
		return false
	}
	if patternCfg1.IgnoreCase != patternCfg2.IgnoreCase {
		return false
	}
	//TODO: add pattern sorting and equality comparison
	return true
}

func isFileCfgEqual(fileCfg1 *FileCfg, fileCfg2 *FileCfg) bool {
	//TODO: find equality between file configs
	return true
}

func isCfgEqual(config1 *Cfg, config2 *Cfg) bool {
	return config1.IsTerm == config2.IsTerm &&
		isFileCfgEqual(&config1.FileCfg, &config2.FileCfg) &&
		isPatternCfgEqual(&config1.PatternCfg, &config2.PatternCfg) &&
		isDisplayCfgEqual(&config1.DisplayCfg, &config2.DisplayCfg)
}

func isDisplayCfgEqual(config1 *DisplayCfg, config2 *DisplayCfg) bool {
	return *config1 == *config2
}

var NoArgsError error = errors.New("No args supplied to program")

type PatternList struct {
	patterns []string
}

func (p *PatternList) String() string {
	return fmt.Sprintf("%v", p.patterns)
}

func (p *PatternList) Set(rawValue string) error {
	p.patterns = append(p.patterns, rawValue)
	return nil
}

type IncludeFileList struct {
	file_patterns []string
}

type ExcludeFileList struct {
	file_patterns []string
}

func (p *IncludeFileList) String() string {
	return fmt.Sprintf("%v", p.file_patterns)
}

func (p *IncludeFileList) Set(rawValue string) error {
	p.file_patterns = append(p.file_patterns, rawValue)
	return nil
}

func ParseOptions(args []string, flagSet *flag.FlagSet) (*Cfg, error) {
	config := &Cfg{}
	config.IsTerm = term.IsTerminal(int(os.Stdin.Fd()))

	flagSet.IntVar(&config.BeforeContext, BEFORE_CONTEXT, 0, "Number of lines above a match to print. Will be overwritten by the -C flag")
	flagSet.IntVar(&config.AfterContext, AFTER_CONTEXT, 0, "Number of lines below a match to print. Will be overwritten by the -C flag")
	flagSet.Int(CONTEXT, 0, "print n lines above and below a given match. This value will override the -A and -B flags")
	flagSet.BoolVar(&config.WordRegexp, WORD_REGEXP, false, "Only return pattern matches that are full words. TODO: define full word.")
	flagSet.BoolVar(&config.WordRegexp, WORD_REGEXP_ALIAS, false, "Alias for --word-regexp.")
	flagSet.BoolVar(&config.LineRegexp, LINE_REGEXP, false, "Only return pattern matches that take up the whole line, not including trailing new line.")
	flagSet.BoolVar(&config.LineRegexp, LINE_REGEXP_ALIAS, false, "Alias for --line-regexp.")
	color := flagSet.String(COLOR, "", "set colorizing for text matches in the GREP_COLOR_2 env var. auto=colorize for terminal output but not for pipe output. 'never' turns off colorization. 'always' forces colorization to any output")

	flagSet.BoolVar(&config.UseCount, COUNT, false, "show only a count of matching lines")
	flagSet.IntVar(&config.BeforeContext, BEFORE_ALIAS, 0, "alias for --before-context")
	flagSet.IntVar(&config.AfterContext, AFTER_ALIAS, 0, "alias for --after-context")
	flagSet.Int(CONTEXT_ALIAS, 0, "alias for --context")
	patterns := &PatternList{}
	flagSet.Var(patterns, REGEXP, "Adds a new pattern to search for. Allows providing more than one pattern to search. Adding multiple --regexp to flags will OR patterns.")
	flagSet.Var(patterns, REGEXP_ALIAS, "alias for --regexp")
	flagSet.StringVar(color, COLOUR, *color, "alias for --color")
	flagSet.BoolVar(&config.UseCount, COUNT_ALIAS, config.UseCount, "alias for --count")

	if err := flagSet.Parse(args); err != nil {
		return nil, err
	}
	//Do not carry over flag.Value interface to CLI config
	config.patterns = patterns.patterns
	// generate the --help config before erroring so we can print the usage guide
	if len(args) == 0 {
		return nil, NoArgsError
	}

	patternArg := flagSet.Arg(0)
	if patternArg != "" {
		config.patterns = append(config.patterns, patternArg)
	}

	if len(config.patterns) == 0 {
		return nil, NoArgsError
	}

	var validationError error
	flagSet.Visit(func(f *flag.Flag) {
		if slices.Contains(ContextAliases, f.Name) {
			contextInt, err := strconv.Atoi(f.Value.String())
			if err != nil {
				validationError = errors.New("Unable to parse context lines from --context into int")
				return
			}
			config.AfterContext = contextInt
			config.BeforeContext = contextInt
		}
		if slices.Contains(ColorAliases, f.Name) {
			if !IsColorOption(f.Value.String()) {
				validationError = fmt.Errorf("Unknown option for %s flag", f.Name)
				return
			}
			config.Color = ColorOption(*color)
		}
	})

	if validationError != nil {
		return nil, validationError
	}

	return config, nil
}
