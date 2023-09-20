BINARY_NAME = utlgo

.PHONY: build-image
build-image:
	docker build -t utlgo:$$(git rev-parse --short HEAD) -t utlgo:latest . 

build-local-binary:
	go build -o ~/.local/bin/$(BINARY_NAME) -trimpath main.go
	chmod u+x ~/.local/bin/$(BINARY_NAME)

