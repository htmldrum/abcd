all:
	go run make.go -v
abcd:
	go run make.go -v -targets abcd
abcv:
	go run make.go -v -targets abcv
test_abcd:
	go test github.com/htmldrum/abcd/cmd/abcd
test_abcv:
	go test github.com/htmldrum/abcd/cmd/abcv
test:
	test_abcd
	test_abcv


