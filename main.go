package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	color := flag.String("color", "", "set a color for text matches in the GREP_COLOR_2 env var")
	flag.StringVar(color, "colour", *color, "alias for --color")
	count := flag.Bool("count", false, "show only a count of matching lines")
	flag.BoolVar(count, "c", *count, "alias for --count")
	flag.Parse()
	
	//what happens when no args are supplied
	//0 positional args = exit with help message
	//1 arg = only pattern, allow reading input from stdin
	// >2 args = both pattern and path, attempt path parsing on args from index 1 and onward since index 0 is the pattern

	args := flag.Args()
	if len(args) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	pattern := args[0]
	_ = pattern
	if len(args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			if scanner.Scan() {
				fmt.Printf("Hello, %s!\n", scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
		}
	}

	for arg_idx := 2; arg_idx < len(args); arg_idx++ {
		// pathArg := args[arg_idx]
		// fmt.Printf("path: %v", pathArg)
		// TODO: file path processing
	}

}
