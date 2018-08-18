get-tests:
	git clone git@github.com:rdebath/Brainfuck ./bfft
	cp -r bfft/testing bf-tests
	rm -rf bfft
test:
	go test
build-test:
	go build && ./go-fuck
