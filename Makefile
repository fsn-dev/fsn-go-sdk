.PHONY: all test clean distclean fmt
.PHONY: account ethkey rlpdump
.PHONY: fsn-cli mongosync

# to prevent mistakely run 'bash Makefile',
ifneq ($(OUTPUT_DIR),)
endif

all:
	./scripts/build.sh $(shell ls -F | grep /$$)

fsn-cli: 
	./scripts/build.sh fsn-cli

mongosync:
	./scripts/build.sh mongosync

account:
	./scripts/build.sh account

ethkey:
	./scripts/build.sh ethkey

rlpdump:
	./scripts/build.sh rlpdump

test:
	@echo "testing Done"

clean:
	go clean -cache
	rm -rf bin

distclean:
	go clean -cache
	rm -rf bin vendor

fmt:
	./scripts/gofmt.sh
