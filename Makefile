sel=.	# selection for test/bench
cnt=5

go:
	go fmt 		./...
	go test		. ./color
	go install	./cmd/fractx

bench:
	go test . -bench=$(sel) -count $(cnt) -benchmem

other:
	GOARCH=386 go build -o fractx386 ./cmd/fractx

.PHONY: go bench other
