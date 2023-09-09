build:
	CGO_ENABLED=0 go build -o bin/ci main.go

build-image: build
	docker build . -t deleema1/circle-ci-test:latest

test-image: build-image
	docker run -it deleema1/circle-ci-test:latest

test:
	go test -v ./...