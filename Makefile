LINT_COMMAND = golangci-lint run

.PHONY: test
test:
	gotestsum \
		--format short-verbose \
		--packages="./..." \
		--junitfile TEST-$*.xml

.PHONY: lint
lint:
	$(LINT_COMMAND)

.PHONY: lint-fix
lint-fix:
	$(LINT_COMMAND) --fix

.PHONY: install
install:
	go mod download
	go install gotest.tools/gotestsum
	go install github.com/golangci/golangci-lint/cmd/golangci-lint
	go install mvdan.cc/gofumpt@latest
