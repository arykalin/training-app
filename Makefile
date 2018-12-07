default: build build-docker
build:
	@echo "building"
	@echo "GOPATH=${GOPATH}"
	env CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -ldflags '-s -w -extldflags "-static"' -a .

build-docker:
	docker build -t arykalin/training-app:latest .

push:
	docker push arykalin/training-app:latest