package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"slices"

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
	IsTerm        bool
	Color         ColorOption
	IgnoreCase    bool
	WordRegexp    bool
	LineRegexp    bool
	ContextLines  [2]int
	UseByteOffset bool
	UseCount      bool
	Patterns      []string
}

var NoArgsError error = errors.New("No args supplied to program")

func parseOptions(args []string) (*cfg, error) {
	if len(args) == 0 {
		return nil, NoArgsError
	}

	config := &cfg{}
	flagSet := flag.NewFlagSet("customgrep", flag.ExitOnError)
	config.IsTerm = term.IsTerminal(int(os.Stdin.Fd()))

	var contextRange [2]int
	linesAbove := flagSet.Int(BEFORE_CONTEXT, 0, "Number of lines above a match to print. Will be overwritten by the -C flag")
	linesBelow := flagSet.Int(AFTER_CONTEXT, 0, "Number of lines below a match to print. Will be overwritten by the -C flag")
	context := flagSet.Int(CONTEXT, 0, "print n lines above and below a given match. This value will override the -A and -B flags")
	color := flagSet.String(COLOR, "", "set colorizing for text matches in the GREP_COLOR_2 env var. auto=colorize for terminal output but not for pipe output. 'never' turns off colorization. 'always' forces colorization to any output")
	//TODO: add flags with custom setters, like -e
	flagSet.BoolVar(&config.UseCount, COUNT, false, "show only a count of matching lines")
	flagSet.Int(BEFORE_ALIAS, 0, "alias for --before-context")
	flagSet.Int(AFTER_ALIAS, 0, "alias for --after-context")
	flagSet.Int(CONTEXT_ALIAS, 0, "alias for --context")

	flagSet.StringVar(color, COLOUR, *color, "alias for --color")
	flagSet.BoolVar(&config.UseCount, COUNT_ALIAS, config.UseCount, "alias for --count")

	contextRange[0] = *linesAbove
	contextRange[1] = *linesBelow

	var validationError error
	flag.Visit(func(f *flag.Flag) {
		if slices.Contains(contextAliases, f.Name) {
			contextRange[0] = *context
			contextRange[1] = *context
		}
		if slices.Contains(colorAliases, f.Name) {
			if !isColorOption(f.Name) {
				validationError = fmt.Errorf("Unknown option for %s flag", f.Name)
			}
		}
	})
	if err := flagSet.Parse(args); err != nil || validationError != nil {
		return nil, err
	}

	return config, nil
}

func main() {
	_, err := parseOptions(os.Args[1:])
	if err != nil {
		if errors.Is(err, NoArgsError) {
			flag.PrintDefaults()
			os.Exit(2)
		}
	}

	//what happens when no args are supplied

	//0 positional args = exit with help message
	args := flag.Args()
	if len(os.Args) == 1 {
		flag.PrintDefaults()
		os.Exit(2)
	}

	//1 arg = only pattern, allow reading input from stdin
	pattern := args[0]
	_ = pattern
	if len(args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			if scanner.Scan() {
				input := scanner.Text()
				fmt.Printf("Hello, %s!\n", input)
				//TODO: process regex, use flags accordingly(no flag is strict string match)
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
		}
	}

	// >=2 args = both pattern and path(s), attempt path parsing on args from index 1 and onward since index 0 is the pattern
	for arg_idx := 1; arg_idx < len(args); arg_idx++ {
		pathArg := args[arg_idx]
		fmt.Printf("path: %v\n", pathArg)
		// TODO: file path processing
	}
}

//TODO: allow parsing of regex with different flags
//TODO:
//TODO:
//TODO:
//TODO:
//TODO:
