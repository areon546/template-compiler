run:
	go build
	./template-compiler -t content



t:
	go test -cover
	./template-compiler -t test -c test -o test-out

log:
	make > attemptLog.lg

tree:
	tree templates/ content/ docs/


coverage:
	go test -coverprofile cover.out && go tool cover -html=cover.out
