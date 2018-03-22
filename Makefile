all: build run

.PHONY: build
build: 
	go build -o main main.go

.PHONY: run
run:
	./main

.PHONY: docker
docker:
	$(shell docker run -p 80:80 go-graphql-todo)

.PHONY: setup_db
setup_db:
	$(shell bin/setup_db.sh)
