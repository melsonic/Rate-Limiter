run:
	go run main.go

build:
	go build -o rl main.go

buildnrun: build
	./rl

test_token_bucket:
	k6 run ./algo/tests/test_token_bucket.js

test_fixed_window:
	k6 run ./algo/tests/test_fixed_window.js

test_sliding_window_log:
	k6 run ./algo/tests/test_sliding_window_log.js

test_sliding_window_counter:
	k6 run ./algo/tests/test_sliding_window_counter.js

clean:
	rm ./rl
