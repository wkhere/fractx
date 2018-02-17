sel=.	# selection for test/bench

go:
	go fmt
	go build # needed because Exec is also tested
	go test
	go install

bench:
	go build
	go test -bench=$(sel)

other:
	GOARCH=386 go build -o fractx386

.PHONY: go bench other
