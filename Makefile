all:
	make darwin
	make linux


linux:
	GOOS=linux GOARCH=amd64 go build -o colorlog.linux

darwin:
	GOOS=darwin GOARCH=amd64 go build -o colorlog

prepare:
	go get -u github.com/alecthomas/participle

clean:
	go clean
	rm -f colorlog.linux colorlog
