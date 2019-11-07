.PHONY: all test clean distclean fmt
.PHONY: vendor vendor_with_proxy
.PHONY: account ethkey rlpdump
.PHONY: fsn-cli mongosync

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

bin/%:
	./scripts/build.sh $(notdir $@)

test:
	@echo "testing Done"

clean:
	go clean -cache
	rm -rf bin

distclean:
	go clean -cache
	rm -rf bin vendor go.sum

vendor:
	./scripts/gomod.sh

vendor_with_proxy:
	./scripts/gomod.sh --proxy

fmt:
	./scripts/gofmt.sh
