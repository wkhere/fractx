PROG=$(shell basename `pwd`)

main: result check

result: go
	./$(PROG)

go:
	go fmt
	go vet
	go build
	@#go test

check:
	md5sum -c MD5

bench:
	go test -bench=.

other:
	GOARCH=386 go build -o fractx386

.PHONY: main go result check bench other
