.PHONY: all test list build clean

ARTIFACT := \
			url2j.linux-amd64 \
			url2j.linux-arm64 \
			url2j.darwin-amd64 \
			#url2.darwin-arm64

all: test build

test:
	go test ./...

list:
	go tool dist list

build: $(ARTIFACT)

url2j.%:
	GOOS=$$(echo $* | cut -d"-" -f1) GOARCH=$$(echo $* | cut -d"-" -f2) \
		 go build -o $@

clean:
	rm -rf $(ARTIFACT)
