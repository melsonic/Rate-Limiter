run:
	go run main.go

build:
	go build -o rl main.go

buildnrun: build
	./rl

test_token_bucket:
	k6 run ./algo/test_token_bucket.js

test_fixed_window:
	k6 run ./algo/test_fixed_window.js

clean:
	rm ./rl
