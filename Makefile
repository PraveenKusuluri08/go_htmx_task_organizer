GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)

default: tidy build htmx_generate run

tidy:
	go mod tidy

build:
	cd cmd && go build -o htmx_app main.go

htmx_generate:
	pwd && templ generate

run:
	cd cmd && air