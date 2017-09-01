all: test

build:
	go build

test: build
	go test -v ./...

integrate: test
	gosearch glide | grep -q 'github.com/Masterminds/glide'
	gosearch echo | grep -q 'github.com/labstack/echo'
	gosearch -c logging | grep -q logrus
	gosearch -c all | grep -q logging
	gosearch -r | grep -q '## Logging'

install: build
	go install .

ci: install integrate

data:
	./updatedata.sh
.PHONY: data
