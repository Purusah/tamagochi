test:
	go test -v -covermode=atomic -coverprofile=coverage.out
	go tool cover -html=coverage.out

lint:
	golangci-lint run -v