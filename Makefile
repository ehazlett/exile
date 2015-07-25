CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
PROJECT=exile
REPO=ehazlett/$(PROJECT)
TAG=${TAG:-latest}
OS="darwin windows linux"
ARCH="amd64 386"
COMMIT=`git rev-parse --short HEAD`

all: build

clean:
	@rm -rf $(PROJECT) $(PROJECT)_*

build:
	@godep go build -a -tags 'netgo' -ldflags "-w -X github.com/ehazlett/exile/version.GitCommit $(COMMIT) -linkmode external -extldflags -static" .

image: build
	@echo Building image $(TAG)
	@docker build -t $(REPO):$(TAG) .

release: deps build image
	@docker push $(REPO):$(TAG)

test:
	@bats test/integration/cli.bats test/integration/certs.bats

.PHONY: all build clean image test release
