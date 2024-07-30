GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)

default: tidy htmx_generate build run

tidy:
	go mod tidy

htmx_generate:
	pwd && templ generate

build:
	cd cmd && go build -o htmx_app main.go

run:
	cd cmd && air