run:
	go run .

watch:
	find . -name '*.go' | entr -r make run
