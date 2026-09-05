package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	optionparse "grep.coding.com/optionParse"
)

func main() {
	flagSet := flag.NewFlagSet("customgrep", flag.ExitOnError)
	_, err := optionparse.ParseOptions(os.Args[1:], flagSet)
	if err != nil {
		if errors.Is(err, optionparse.NoArgsError) {
			flagSet.Usage()
			os.Exit(2)
		}
		fmt.Println(err.Error())
		os.Exit(2)
	}

	//1 arg = only pattern, allow reading input from stdin
}
