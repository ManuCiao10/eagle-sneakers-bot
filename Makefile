all:
	rm -rf EagleBot_0.0.24.exe
	go generate
	env GOOS=windows GOARCH=amd64 go build -o EagleBot_0.0.24.exe github.com/eagle

git:
	@git add .
	@read -p "Insert commit name: " TAG && git commit -m "$$TAG"
	@git push

# 64-bit
macOs:
	rm -rf EagleBot_0.0.24
	go generate
	rm resource.syso
	env GOOS=darwin GOARCH=amd64 go build github.com/eagle
