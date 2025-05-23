run:
	go build && go run .

test:
	go test -cover

coverage:
	go test -coverprofile cover.out && go tool cover -html=cover.out
