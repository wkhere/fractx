sel=.	# selection for test/bench/fuzz
opt=	# options for fuzz
cnt=5

go:
	go vet 		./...
	go test		. ./color
	go build	./cmd/fractx

install:
	go install	./cmd/fractx

fuzz:
	#go test -fuzz=$(sel) $(opt) .
	go test -fuzz=$(sel) $(opt) ./color

bench:
	go test . -bench=$(sel) -count $(cnt) -benchmem

cover:
	go test -coverprofile cov . ./color
	go tool cover -html cov

other:
	GOARCH=386 go build -o fractx386 ./cmd/fractx

.PHONY: go fuzz bench cover other
