# Parts of the Makefile are adapted from the OpenTofu project, licensed under
# the MPL-2.0.
# See: https://github.com/opentofu/opentofu/blob/main/Makefile
LICENSEI_VERSION = 0.9.0

.PHONY: build
build:
	go build -o ager ./main.go

.PHONY: golangci-lint
golangci-lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.2 run --timeout 60m ./...

.PHONY: license-check
license-check:
	go mod vendor
	bin/licensei cache --debug
	bin/licensei check --debug
	bin/licensei header --debug
	rm -rf vendor/
	git diff --exit-code

deps: bin/licensei
deps:

bin/licensei: bin/licensei-${LICENSEI_VERSION}
	@ln -sf licensei-${LICENSEI_VERSION} bin/licensei
bin/licensei-${LICENSEI_VERSION}:
	@mkdir -p bin
	curl -sfL https://git.io/licensei | bash -s v${LICENSEI_VERSION}
	@mv bin/licensei $@
