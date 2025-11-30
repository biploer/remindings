BINPATH=bin/remindings

.PHONY: build
build: build-app

.PHONY: build-app
build-app:
	go build -o $(BINPATH) cmd/remindings/main.go

.PHONY: run
run: build
	$(BINPATH)
