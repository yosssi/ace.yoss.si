bin:
	go-bindata -o asset.go public/... views/...
run:
	PORT=8080 go run main.go asset.go
