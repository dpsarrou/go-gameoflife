test:
	@go test -cover ./gameoflife/...

test-coverage:
	@go test -cover -coverprofile=./tmp/coverage.out ./gameoflife/...
	@go tool cover -html=./tmp/coverage.out

run:
	@go run .

run-glider:
	@go run . -glider