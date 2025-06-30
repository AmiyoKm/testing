test:
	go test -v ./...

test-bench:
	go test -bench=. ./... -benchmem

test-all: test test-bench

test-d:
	cd ${d} && go test -v

test-cover:
	go test -cover ./...