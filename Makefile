sel=.	# selection for test/bench/fuzz
opt=	# options for fuzz
cnt=5

go:
	go fmt 		./...
	go test		. ./color
	go install	./cmd/fractx

fuzz:
	#go test -fuzz=$(sel) $(opt) .
	go test -fuzz=$(sel) $(opt) ./color

bench:
	go test . -bench=$(sel) -count $(cnt) -benchmem

other:
	GOARCH=386 go build -o fractx386 ./cmd/fractx

.PHONY: go bench other
