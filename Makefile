build:
	go build -mod vendor -work -v -o bin/metaparser cmd/metaparser/metaparser.go

SOURCE = $(wildcard *.go)
TAG ?= $(shell git describe --tags)
GOBUILD = go build -ldflags '-w'

ALL = $(foreach suffix,win.exe linux osx,\
		bin/metaparser-$(suffix))

all: $(ALL)

clean:
	rm -f $(ALL)

test:
	go test
	cram tests/cram.t

win.exe = windows
osx = darwin
bin/metaparser-%: $(SOURCE)
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=$(firstword $($*) $*) GOARCH=amd64 go build -o $@

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
