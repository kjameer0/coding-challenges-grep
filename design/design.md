# My Design

## Regex handling

It appears that grep uses an older model of regular expressions that might allow for exponential time searches. This [article](https://owasp.org/www-community/attacks/Regular_expression_Denial_of_Service_-_ReDoS) goes into how ReDoS(Regex Denial of Service) attacks can allow attackers to misuse regex that backtracks to do advanced pattern matching and cause a system to slow down dramatically. Technically I don't really have to worry about this that much since the user is only using the CLI to interact with the program. If there was a web component to this, I would feel differently. That being said, I think I can build this with multiple regexes in mind, one of which is go's regex library.

Basically I need to be able to provide a single line and an option of which regex to use, and then match the line with the regex strategy. I can let the library give out the allowed regexes.

Here is an example use.

```go
  line := "hello"
  patterns := []string{"ell*"}
  //
  matches := regexPackage.SearchLine(line, patterns, regexPackage.ExtendedRegex)
```

This example is not that configurable from the user side. But what would a user need to configure? We might want to include things like ignore case. We might also want to pass negation so only non-matches come back. So basically we can have a config that lets you set ignore case and other options, and if the regex strategy supports it, let people opt in. Not all regex types support the same flags? I guess all support ignore case.

## Labeling

There are multiple kinds of labels:

1. file name
2. line number
3. column number (index in line)
4. byte offset in file (where first byte in file is byte offset 0)
