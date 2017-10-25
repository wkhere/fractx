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

.PHONY: main go result check
