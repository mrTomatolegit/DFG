VERSION=0.2.0

OSFLAG :=
ifeq ($(OS),Windows_NT)
	OSFLAG += windows
else
	OSFLAG += dfg
endif

build:
	go build -o ./bin/dfg.exe ./cmd/${OSFLAG}/dfg.go

run:
	go run main.go

compile:
	@echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/dfg_amd64_${VERSION}_darwin ./cmd/dfg/dfg.go
	GOOS=linux GOARCH=amd64 go build -o bin/dfg_amd64-${VERSION}_linux ./cmd/dfg/dfg.go
	GOOS=windows GOARCH=amd64 go build -o bin/dfg_amd64_${VERSION}_windows.exe ./cmd/windows/dfg.go

all: build
