PROG=$(shell basename `pwd`)

main: result check

result: go
	./$(PROG)

go:
	go fmt
	go vet
	go build

check:
	md5sum -c MD5

bench:
	go test -bench=.

.PHONY: main go result check bench
