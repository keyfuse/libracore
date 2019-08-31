# libracore
export PATH := $(GOPATH)/bin:$(PATH)

clean:
	@echo "--> Cleaning..."
	@go clean
	@rm -f bin/*

test:
	@echo "--> Testing..."
	go get -v github.com/stretchr/testify/assert
	@$(MAKE) testxbase
	@$(MAKE) testxcrypto
	@$(MAKE) testxcore

testxbase:
	go test -v ./xbase

testxcrypto:
	go test -v ./xcrypto/...

testxcore:
	go test -v ./xcore

runexamples:
	go run examples/address.go
	go run examples/account_mint.go
	go run examples/account_query.go
	go run examples/p2p_transaction.go
	go run examples/p2p_transfer.go
	go run examples/two_party_ed25519_transaction.go

# code coverage
allpkgs =	./xbase ./xcrypto/... ./xcore
coverage:
	go get -v github.com/pierrre/gotestcover
	gotestcover -coverprofile=coverage.out -v $(allpkgs)
	go tool cover -html=coverage.out

check:
	go get -v gopkg.in/alecthomas/gometalinter.v2
	$(GOPATH)/bin/gometalinter.v2 -j 4 --disable-all \
	--enable=gofmt \
	--enable=golint \
	--enable=vet \
	--enable=gosimple \
	--enable=unconvert \
	--deadline=10m $(allpkgs) 2>&1 | tee /dev/stderr

.PHONY: clean test coverage check
