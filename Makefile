PACKAGES=$(shell go list ./... | grep -v /vendor/)
RACE := $(shell test $$(go env GOARCH) != "amd64" || (echo "-race"))

# Program version
VERSION := $(shell grep "const Version " server/version.go | sed -E 's/.*"(.+)"$$/\1/')

help:
	@echo 'Available commands:'
	@echo
	@echo 'Usage:'
	@echo '    make deps     Install go deps.'
	@echo '    make build    Compile the project.'
	@echo '    make restore  Restore all dependencies.'
	@echo '    make clean    Clean the directory tree.'
	@echo

test: ## run tests, except integration tests
	@go test ${RACE} ${PACKAGES}

deps:
	go get github.com/tcnksm/ghr
	go get github.com/mitchellh/gox

build:
	@echo "Compiling..."
	@mkdir -p ./bin
	@gox -output "bin/{{.Dir}}_${VERSION}_{{.OS}}_{{.Arch}}" -os="linux" -os="darwin" -arch="386" -arch="amd64" ./
	@go build -i -o ./bin/ec2-metadata
	@echo "All done! The binaries is in ./bin let's have fun!"

build/docker: build
	@docker build -t pixelmilk/ec2-metadata:latest .

vet: ## run go vet
	@test -z "$$(go vet ${PACKAGES} 2>&1 | grep -v '*composite literal uses unkeyed fields|exit status 0)' | tee /dev/stderr)"

ci: vet test

restore:
	@dep ensure
