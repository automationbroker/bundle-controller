vendor:
	dep ensure

compile:
	go build -i -ldflags="-s -w" ./cmd/main.go

run: compile
	@./main
