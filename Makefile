BINPATH=bin/remindings
IMAGE=remindings:local

.PHONY: build
build: build-app

.PHONY: build-app
build-app:
	go build -o $(BINPATH) cmd/remindings/main.go

.PHONY: run
run: build
	$(BINPATH)

.PHONY: dock
dock: dock-build
	podman run -it --rm localhost/$(IMAGE)

.PHONY: dock-build
dock-build:
	podman build -t $(IMAGE) .
