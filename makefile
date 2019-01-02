NAME := truffle
PKG := github.com/JoelPagliuca/$(NAME)

GITCOMMIT := $(shell git rev-parse --short HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
	GITCOMMIT := $(GITCOMMIT)-dirty
endif

GO_LDFLAGS=-ldflags "-w"

BUILDDIR=build

.PHONY: help
help: ## Print this message and exit
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


.PHONY: build
build: $(NAME) ## Builds an executable for host platform

$(NAME): $(wildcard *.go)
	@echo "+ $@"
	go build $(GO_LDFLAGS) -o $(NAME) .

.PHONY: build-linux
build-linux: *.go ## linux compilation
	@echo "+ $@"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GO_LDFLAGS) -o $(BUILDDIR)/$(NAME)-linux
	sha256sum $(BUILDDIR)/$(NAME)-linux > $(BUILDDIR)/$(NAME)-linux.sha256

.PHONY: build-macos
build-macos: *.go ## macos compilation
	@echo "+ $@"
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(GO_LDFLAGS) -o $(BUILDDIR)/$(NAME)-macos
	sha256sum $(BUILDDIR)/$(NAME)-macos > $(BUILDDIR)/$(NAME)-macos.md5;

build-windows: *.go ## windows compilation
	@echo "+ $@"
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(GO_LDFLAGS) -o $(BUILDDIR)/$(NAME).exe
	sha256sum $(BUILDDIR)/$(NAME).exe > $(BUILDDIR)/$(NAME).exe.md5;

.PHONY: cross
cross: ## Builds the cross-compiled binaries
	@echo "+ $@"
	mkdir -p $(BUILDDIR)
	make build-linux
	make build-macos
	make build-windows

test-setup: build ## Setup the test project for testing
	@echo "+ $@"
	rm -rf test-project
	./scripts/setup-test-harness.sh

test: test-setup ## Run some tests against a test git project
	@echo "+ $@"
	./truffle -i test-project
	./scripts/run-tests.sh

.PHONY: clean
clean: ## Cleanup any build binaries or packages
	@echo "+ $@"
	rm -f $(NAME)
	rm -rf $(BUILDDIR)