GCI_VERSION = 0.13.5
GOFUMPT_VERSION = 0.7.0
GOLANGCI_LINT_VERSION = 1.62.2
LICENSEI_VERSION = 0.9.0

.PHONY: all
all: build

.PHONY: build
build:
ifneq ($(AGRICOLA_VERSION),)
	go build -ldflags "-X 'github.com/anttikivi/agricola/version.buildVersion=$(AGRICOLA_VERSION)'" -o ager ./main.go
else
	go build -o ager ./main.go
endif

.PHONY: fmt
fmt:
	go run github.com/daixiang0/gci@v${GCI_VERSION} print . --skip-generated -s standard -s default
	go run mvdan.cc/gofumpt@v${GOFUMPT_VERSION} -l -w .

.PHONY: test
test:
	go test -v ./...

.PHONY: check
check: lint license-check

.PHONY: lint
lint: golangci-lint

.PHONY: golangci-lint
golangci-lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v${GOLANGCI_LINT_VERSION} run ./...

.PHONY: license-check
license-check:
	go mod vendor
	bin/licensei cache --debug
	bin/licensei check --debug
	bin/licensei header --debug
	rm -rf vendor/
	git diff --exit-code

deps: bin/licensei

bin/licensei: bin/licensei-${LICENSEI_VERSION}
	@ln -sf licensei-${LICENSEI_VERSION} bin/licensei
bin/licensei-${LICENSEI_VERSION}:
	@mkdir -p bin
	curl -sfL https://git.io/licensei | bash -s v${LICENSEI_VERSION}
	@mv bin/licensei $@
