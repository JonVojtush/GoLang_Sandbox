build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go run main.go

run: build
	./OpenNomad_GoApps

clean:
	@go clean
	@-rm web/app.wasm