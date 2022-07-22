build: build-win


.PHONY: build-win
build-win: 
	@rm fyne-demo.exe
	@go mod tidy
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build .