all: build silent-test

run: vendor_get
	 go run asio.go

build: vendor_get
		go build -o bin/asio asio.go

test: vendor_get
	ginkgo -r -v -cover --randomizeAllSpecs

silent-test:
	go test $(PACKAGES)

format:
	go fmt $(PACKAGES)

vendor_get:
	go get -u github.com/codegangsta/negroni \
	github.com/julienschmidt/httprouter \
	github.com/modocache/gory \
	github.com/onsi/ginkgo \
	github.com/onsi/ginkgo/ginkgo \
	github.com/onsi/gomega \
