# File Path Resolution

## Description

For any provided path, I have to determine if it's a file or a dir
What problem am i solving?
I have to take each file path arg and try to resolve it down to the actual set of files we will be choosing from. so I'm turning an array of glob patterns into a set. so if i go and take a file path like ./hi.go, I need the program to be able to detect that that is not a directory and return the set of files as `["hi.go"]`. the shell will expand files passed as arguments, and I will have to process args passed through the include and exclude flags myself


## Requirements
