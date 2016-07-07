abcd:
	go build github.com/htmldrum/abcd/cmd/abcd
abcv:
	go build github.com/htmldrum/abcd/cmd/abcv
test_abcd:
	go test github.com/htmldrum/abcd/cmd/abcd
test_abcv:
	go test github.com/htmldrum/abcd/cmd/abcv
test:
	test_abcd
	test_abcv
