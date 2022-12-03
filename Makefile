all:
	rm -rf EagleBot_0.0.23.exe
	go generate
	env GOOS=windows GOARCH=amd64 go build -o EagleBot_0.0.23.exe github.com/eagle


	