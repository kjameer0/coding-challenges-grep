package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("hell")
	color := flag.String("color", "", "set a color for text matches in the GREP_COLOR_2 env var")
	flag.StringVar(color, "colour", *color, "alias for --color")
	count := flag.Bool("count", false, "show only a count of matching lines")
	flag.BoolVar(count, "c", *count, "alias for --count")
	flag.Parse()
	
}
