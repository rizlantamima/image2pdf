build:
	go build -o bin/cli cmd/cli/main.go

run: build
	./bin/cli $(ARGS)


compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o bin/cli-linux-amd64 cmd/cli/main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/cli-darwin-arm64 cmd/cli/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/cli-windows-amd64 cmd/cli/main.go