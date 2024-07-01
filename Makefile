.PHONY=all clean build test

all: clean build test

clean:
	rm -f main main.db models.go *.sql.go

build:
	sqlc generate
	go build ./...

test:
	./main
