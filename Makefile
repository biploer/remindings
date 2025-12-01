BINPATH=bin/remindings
IMAGE=remindings:local

.PHONY: build
build: test build-app

.PHONY: build-app
build-app:
	go build -o $(BINPATH) cmd/remindings/main.go

.PHONY: test
test:
	go test ./...

.PHONY: run
run: build
	$(BINPATH)

.PHONY: dock
dock: test dock-build
	podman run -it --rm localhost/$(IMAGE)

.PHONY: dock-build
dock-build:
	podman build -t $(IMAGE) .
