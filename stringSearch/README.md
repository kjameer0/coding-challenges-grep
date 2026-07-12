# String Searching Design

## Interface

The basic idea is that there should be multiple strategies to pick from to search with. Examples include:

1. fixed string search
2. RE2 regex search

A caller needs to provide a line of text to search, as well as all of the patterns they are looking for, and the package will return a list of strings that match, as well as their locations in the line.

## Line Locations

For any line, there are zero or more characters. A line can have different encodings as well. This package will assume strings are UTF-8, another package should handle text conversion.

## Buffering long lines

Long lines should be chunked so that we don't load 1GB long lines. This can be a hard limit introduced after a certain size.


