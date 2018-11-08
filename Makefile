all:
	make darwin
	make linux

linux:
	GOOS=linux GOARCH=amd64 go build -o colorlog

darwin:
	GOOS=darwin GOARCH=amd64 go build -o colorlog.darwin

prepare:


clean:
	go clean
	rm -f colorlog.linux colorlog
