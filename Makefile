run:
	go build
	make c

t:
	go test -cover
	make c

c:
	template-compiler -t test -c test -o test-out

bin:
	go build
	mv ./template-compiler ~/.local/bin/template-compiler



log:
	make > attemptLog.lg

tree:
	tree content/ docs/


coverage:
	go test -coverprofile cover.out && go tool cover -html=cover.out
