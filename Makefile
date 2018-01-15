all: build run

.PHONY: build
build: 
	go build -o main main.go

.PHONY: run
run:
	./main

.PHONY: setup_db
setup_db:
	$(shell bin/setup_db.sh)
