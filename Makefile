.PHONY=all clean build test

all: clean build test

clean:
	rm -f main main.db *.sql.go

build:
	sqlc generate
	go build ./...

test:
	./main
