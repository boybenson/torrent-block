# Build the project
build:
	go build -o myapp main.go

run:
	go run main.go

seed:
	go run main.go -mode=seeder -file=learn.txt -addr=:9000

leech:
	go run main.go -mode=leech -addr=127.0.0.1:9000