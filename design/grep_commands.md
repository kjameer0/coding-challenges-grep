# grep interface / error-handling probe commands

Run from inside `design/`:

```bash
cd /Users/khalidjameer/projects/coding-challenges-grep/design
```

## Basic interface

```bash
grep ; echo $?
grep "Grep" ; echo $?
echo "Grep's Design" | grep "Grep" ; echo $?
```

## Positional arg counts

```bash
grep "Grep" original_grep_design.md ; echo $?
grep "Grep" original_grep_design.md grep_design.excalidraw ; echo $?
grep "Grep" *.md *.excalidraw ; echo $?
```

## Exit code semantics

```bash
grep "Grep" original_grep_design.md ; echo $?
grep "zzz_nomatch" original_grep_design.md ; echo $?
grep "Grep" nonexistent_file.md ; echo $?
```

## Error handling with multiple files, one missing

```bash
grep "Grep" nonexistent_file.md original_grep_design.md ; echo $?
```

## Missing pattern / malformed regex

```bash
grep -- "" original_grep_design.md ; echo $?
grep "[" original_grep_design.md ; echo $?
grep -c "Grep" original_grep_design.md ; echo $?
grep --bogus-flag "Grep" original_grep_design.md ; echo $?
```
- `--` does nothing, unless the pattern involved contains a "-" at the front
-
## Directory as input (no -r)

```bash
grep "Grep" design/ ; echo $?
grep -r "Grep" design/ ; echo $?

```
first command:
- simply errors with "grep: design/: Is a directory"
- exits with 2(genuine use error)
- it is a use error to not use -r and provide a directory
- adding a valid path as an arg after a dir without -r still yields exit code 2, even with actual matches comeback
- this means that if anything errors, the exit code errors(greedy)

second command:
- returns a list of pathnames: the-matching-line
- exits with 0

## Flag/value edge cases

```bash
grep -c ; echo $?
grep --color=always "Grep" original_grep_design.md | cat -v ; echo $?
```
