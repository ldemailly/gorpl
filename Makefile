all: generate lint tests failing-tests

GO_BUILD_TAGS:=no_net,no_json

build:
	CGO_ENABLE=0 go build -trimpath -ldflags="-w -s" -tags "$(GO_BUILD_TAGS)" .

tests:
	go test -race ./...

failing-tests:
	-go test -v ./... -tags=runfailingtests -run TestLetStatementsFormerlyCrashingNowFailingOnPurpose

test: tests

generate:
	go generate ./... # if this fails go install golang.org/x/tools/cmd/stringer@latest

lint: .golangci.yml
	golangci-lint run

.golangci.yml: Makefile
	curl -fsS -o .golangci.yml https://raw.githubusercontent.com/fortio/workflows/main/golangci.yml

.PHONY: all lint generate tests test
