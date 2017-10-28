sel=.	# selection for test/bench

go:
	go fmt
	go vet
	go build
	go test
	go install

bench:
	go test -bench=$(sel)

other:
	GOARCH=386 go build -o fractx386

.PHONY: go bench other
