all:
	make darwin
	make linux

linux:
	GOOS=linux GOARCH=amd64 go build -o cl

darwin:
	GOOS=darwin GOARCH=amd64 go build -o cl.osx

prepare:


clean:
	go clean
	rm -f cl.linux cl
