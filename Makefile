abcd:
	go build github.com/htmldrum/abcd/cmd/abcd
abcv:
	go build github.com/htmldrum/abcd/cmd/abcv
test_abcd:
	go test -race -v github.com/htmldrum/abcd/cmd/abcd
test_abcv:
	go test -race -v github.com/htmldrum/abcd/cmd/abcv
test_divs:
	go test -race -v github.com/htmldrum/abcd/divisions
	go test -race -v github.com/htmldrum/abcd/divisions/news
	go test -race -v github.com/htmldrum/abcd/divisions/news/strategies/v0
test:
	go vet
	@make test_abcd
	@make test_abcv
	@make test_divs
doc:
	godoc
