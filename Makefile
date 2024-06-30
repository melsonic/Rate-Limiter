run:
	go run main.go

build:
	go build -o rl main.go

buildnrun:
	go build -o rl main.go && ./rl
