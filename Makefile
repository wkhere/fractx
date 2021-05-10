sel=.	# selection for test/bench
cnt=5

go:
	go fmt
	go test
	go install

bench:
	go test -bench=$(sel) -count $(cnt)

other:
	GOARCH=386 go build -o fractx386

.PHONY: go bench other
