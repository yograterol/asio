all: build silent-test

run: vendor_get
	 go run asio.go

build: vendor_get
		go build -o bin/asio asio.go

test_cover: vendor_get
	go get -u golang.org/x/tools/cmd/cover
	ginkgo -r -v -cover --randomizeAllSpecs

test: vendor_get
	ginkgo -r -v

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
	github.com/onsi/gomega
