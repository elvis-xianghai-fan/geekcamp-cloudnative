export tag=0.1
root:
	export ROOT=github.com/elvis-xianghai-fan/geekcamp-cloudnative

build:
	echo "building httpserver"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t elvfan/httpserver:${tag} .

push: release
	echo "pushing elvfan/httpserver"
	docker push elvfan/httpserver:${tag}
