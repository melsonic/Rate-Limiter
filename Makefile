run:
	go run main.go

build:
	go build -o rl main.go

buildnrun: build
	./rl

test_token_bucket:
	k6 run ./algo/token_bucket_test.js

clean:
	rm ./rl
