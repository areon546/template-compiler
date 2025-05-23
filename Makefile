run:
	go build
	./template-compiler	-o documents

test:
	go test -cover

coverage:
	go test -coverprofile cover.out && go tool cover -html=cover.out
