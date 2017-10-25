PROG=$(shell basename `pwd`)

result: go
	./$(PROG)

go:
	go fmt
	go vet
	go build

check:
	md5sum -c MD5
