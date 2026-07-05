package main

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
	COLOR_NEVER    = "never"
	COLOR_ALWAYS   = "always"
	COLOR_AUTO     = "auto"
	COLOR          = "color"
	COLOUR         = "colour"
	CONTEXT        = "context"
	AFTER_CONTEXT  = "after-context"
	BEFORE_CONTEXT = "before-context"
	COUNT          = "count"
	COUNT_ALIAS    = "c"
	BEFORE_ALIAS   = "B"
	AFTER_ALIAS    = "A"
	CONTEXT_ALIAS  = "C"
	REGEXP         = "regexp"
	REGEXP_ALIAS   = "e"
)

var contextAliases []string = []string{CONTEXT, CONTEXT_ALIAS}
var afterContextAliases []string = []string{AFTER_CONTEXT, AFTER_ALIAS}
var beforeContextAliases []string = []string{BEFORE_CONTEXT, BEFORE_ALIAS}
var colorAliases []string = []string{COLOR, COLOUR}

func isColorOption(option string) bool {
	allowedOptions := []string{COLOR_AUTO, COLOR_ALWAYS, COLOR_NEVER}
	return slices.Contains(allowedOptions, option)
}

type cfg struct {
	IsTerm bool
	fileCfg
	displayCfg
	patternCfg
}

type fileCfg struct {
	//TODO: plan out types for file inclusion ond exclusion
}

type displayCfg struct {
	Color         ColorOption
	BeforeContext int
	AfterContext  int
	UseByteOffset bool
	UseCount      bool
}

type patternCfg struct {
	WordRegexp bool
	LineRegexp bool
	IgnoreCase bool
	patterns   []string
}

func isPatternCfgEqual(patternCfg1 *patternCfg, patternCfg2 *patternCfg) bool {
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

func isFileCfgEqual(fileCfg1 *fileCfg, fileCfg2 *fileCfg) bool {
	//TODO: find equality between file configs
	return true
}

func isCfgEqual(config1 *cfg, config2 *cfg) bool {
	return config1.IsTerm == config2.IsTerm &&
		isFileCfgEqual(&config1.fileCfg, &config2.fileCfg) &&
		isPatternCfgEqual(&config1.patternCfg, &config2.patternCfg) &&
		isDisplayCfgEqual(&config1.displayCfg, &config2.displayCfg)
}

func isDisplayCfgEqual(config1 *displayCfg, config2 *displayCfg) bool {
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


func parseOptions(args []string, flagSet *flag.FlagSet) (*cfg, error) {
	config := &cfg{}
	config.IsTerm = term.IsTerminal(int(os.Stdin.Fd()))

	flagSet.IntVar(&config.BeforeContext, BEFORE_CONTEXT, 0, "Number of lines above a match to print. Will be overwritten by the -C flag")
	flagSet.IntVar(&config.AfterContext, AFTER_CONTEXT, 0, "Number of lines below a match to print. Will be overwritten by the -C flag")
	flagSet.Int(CONTEXT, 0, "print n lines above and below a given match. This value will override the -A and -B flags")
	color := flagSet.String(COLOR, "", "set colorizing for text matches in the GREP_COLOR_2 env var. auto=colorize for terminal output but not for pipe output. 'never' turns off colorization. 'always' forces colorization to any output")
	//TODO: add flags with custom setters, like -e
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
		if slices.Contains(contextAliases, f.Name) {
			contextInt, err := strconv.Atoi(f.Value.String())
			if err != nil {
				validationError = errors.New("Unable to parse context lines from --context into int")
				return
			}
			config.AfterContext = contextInt
			config.BeforeContext = contextInt
		}
		if slices.Contains(colorAliases, f.Name) {
			if !isColorOption(f.Value.String()) {
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

func main() {
	flagSet := flag.NewFlagSet("customgrep", flag.ExitOnError)
	_, err := parseOptions(os.Args[1:], flagSet)
	if err != nil {
		if errors.Is(err, NoArgsError) {
			flagSet.Usage()
			os.Exit(2)
		}
		fmt.Println(err.Error())
		os.Exit(2)
	}

	//1 arg = only pattern, allow reading input from stdin
}

//TODO: allow parsing of regex with different flags
//TODO:
//TODO:
//TODO:
//TODO:
//TODO:
