NAME := nocommit
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
	mkdir -p $(BUILDDIR)/linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(GO_LDFLAGS) -o $(BUILDDIR)/linux/$(NAME)-linux
	sha256sum $(BUILDDIR)/linux/$(NAME)-linux > $(BUILDDIR)/linux/$(NAME)-linux.sha256

.PHONY: build-macos
build-macos: *.go ## macos compilation
	@echo "+ $@"
	mkdir -p $(BUILDDIR)/macos
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(GO_LDFLAGS) -o $(BUILDDIR)/macos/$(NAME)-macos
	sha256sum $(BUILDDIR)/macos/$(NAME)-macos > $(BUILDDIR)/macos/$(NAME)-macos.md5;

.PHONY: cross
cross: build-linux build-macos ## Builds the cross-compiled binaries
	@echo "+ $@"

test: build ## Run some tests against a test project
	@echo "+ $@"
	./scripts/setup-test-harness.sh

.PHONY: clean
clean: ## Cleanup any build binaries or packages
	@echo "+ $@"
	rm -f $(NAME)
	rm -rf $(BUILDDIR)
	rm -rf test-project