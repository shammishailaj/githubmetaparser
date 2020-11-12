SOURCE = $(wildcard cmd/metaparser/*.go)
TAG ?= $(shell git describe --tags)
GOBUILD = go build -ldflags '-w'
GIT_BRANCH = $(shell git branch --show-current)
VERSION = $(shell cat VERSION)-$(GIT_BRANCH)
OS = $(shell go env GOHOSTOS)
ARCH = $(shell go env GOHOSTARCH)

ALL = $(foreach suffix,win.exe linux osx,\
		bin/metaparser-$(VERSION)-$(suffix))

# If it is the "main" branch do not attach branch name suffix
ifeq ($(GIT_BRANCH),main)
		VERSION := $(shell cat VERSION)
endif

# If it is the "main" branch do not attach branch name suffix
ifeq ($(GIT_BRANCH),master)
	VERSION := $(shell cat VERSION)
endif


build:
	$(info GIT_BRANCH=$(GIT_BRANCH))
	$(info VERSION=$(VERSION))
	govvv build -pkg "github.com/shammishailaj/metaparser/internal/app/metaparser/cmd" -mod vendor -work -v -o bin/metaparser-$(VERSION)-$(OS)-$(ARCH) cmd/metaparser/metaparser.go

all: $(ALL)

clean:
	rm -f $(ALL)

test:
	go test
	cram tests/cram.t

win.exe = windows
osx = darwin
bin/metaparser-$(VERSION)-%: $(SOURCE)
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=$(firstword $($*) $*) GOARCH=amd64 govvv build -pkg "github.com/shammishailaj/metaparser/internal/app/metaparser/cmd" -mod vendor -o $@ $(SOURCE)

ifndef desc
release:
	@echo "Push a tag and run this as 'make release desc=tralala'"
else
release: $(ALL)
	github-release release -u shammishailaj -r metaparser -t "$(TAG)" -n "$(TAG)" --description "$(desc)"
	@for x in $(ALL); do \
		echo "Uploading $$x" && \
		github-release upload -u shammishailaj \
                              -r metaparser \
                              -t $(TAG) \
                              -f "$$x" \
                              -n "$$(basename $$x)"; \
	done
endif
