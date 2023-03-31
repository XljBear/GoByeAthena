.PHONY : mac windows linux all mkdir
mac: prepare
	 CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/GoByeAthena_Mac/GoByeAthena main.go

windows: prepare
	go generate
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/GoByeAthena_Windows/GoByeAthena.exe ./

linux: prepare
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/GoByeAthena_Linux/GoByeAthena main.go

all: mac windows linux

prepare:
	mkdir -p ./build/
	mkdir -p ./build/GoByeAthena_Mac
	mkdir -p ./build/GoByeAthena_Windows
	mkdir -p ./build/GoByeAthena_Linux