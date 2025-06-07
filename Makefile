run:
	go build
	./template-compiler

tree:
	tree templates/ content/ docs/


test:
	go test -cover

coverage:
	go test -coverprofile cover.out && go tool cover -html=cover.out
