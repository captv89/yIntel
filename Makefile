run:
	go run main.go

build:
	go build -o bin/yIntel main.go

rpi-build:
	GOOS=linux GOARCH=arm GOARM=7 go build -o bin/yIntel-rpi main.go

ubuntu-build:
	GOOS=linux GOARCH=amd64 go build -o bin/yIntel-ubuntu main.go