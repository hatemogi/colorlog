all:
	make darwin
	make linux

linux:
	GOOS=linux GOARCH=amd64 go build -o colorlog

darwin:
	GOOS=darwin GOARCH=amd64 go build -o colorlog.darwin

prepare:
	go get -u github.com/alecthomas/participle

clean:
	go clean
	rm -f colorlog.linux colorlog
