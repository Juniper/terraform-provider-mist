HOSTNAME=terraform.local
NAMESPACE=local
NAME=mist
BINARY=terraform-provider-${NAME}
VERSION=0.4.0
OS_ARCH=darwin_arm64
GOBIN ?= $(if $(shell go env GOBIN),$(shell go env GOBIN),$(shell go env GOPATH)/bin)
GOLANGCI_LINT_VER=v2.5.0
TESTDIRS += $(sort $(shell go list ./... | grep -Ev '(cmd|test|mock|fake|tools)'))

default: install

compliance:
    # golang.org/x/sys ignored to avoid producing different results on different platforms (x/sys vs. x/sys/unix, etc...)
    # The license details, if it were included in the Third_Party_Code directory, would look something like this, depending
    # on the build platform:

    ### golang.org/x/sys/unix
    #
    #* Name: golang.org/x/sys/unix
    #* Version: v0.17.0
    #* License: [BSD-3-Clause](https://cs.opensource.google/go/x/sys/+/v0.17.0:LICENSE)
    #
    #
    #Copyright (c) 2009 The Go Authors. All rights reserved.
    #
    #Redistribution and use in source and binary forms, with or without
    #modification, are permitted provided that the following conditions are
    #met:
    #
    #   * Redistributions of source code must retain the above copyright
    #notice, this list of conditions and the following disclaimer.
    #   * Redistributions in binary form must reproduce the above
    #copyright notice, this list of conditions and the following disclaimer
    #in the documentation and/or other materials provided with the
    #distribution.
    #   * Neither the name of Google Inc. nor the names of its
    #contributors may be used to endorse or promote products derived from
    #this software without specific prior written permission.
    #
    #THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
    #"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
    #LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
    #A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
    #OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
    #SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
    #LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
    #DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
    #THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
    #(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
    #OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

	go get github.com/chrismarget-j/go-licenses || exit 1 ;\
	go run github.com/chrismarget-j/go-licenses save   --ignore github.com/Juniper --ignore golang.org/x/sys --save_path Third_Party_Code --force ./... || exit 1 ;\
	go run github.com/chrismarget-j/go-licenses report --ignore github.com/Juniper --ignore golang.org/x/sys --template .notices.tpl ./... > Third_Party_Code/NOTICES.md || exit 1 ;\
	go mod tidy ;

doc:
	tfplugindocs generate

docs-check:
	@sh -c "$(CURDIR)/scripts/tfplugindocs.sh"

tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/$(GOLANGCI_LINT_VER)/install.sh | sh -s -- -b $(GOBIN) $(GOLANGCI_LINT_VER)
	go install golang.org/x/vuln/cmd/govulncheck@latest

build:
	GPG_FINGERPRINT=82AD65745D9BAFF7 goreleaser release --clean 

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

lint:
	@test -s $(GOBIN)/golangci-lint || { echo "golangci-lint does not exist! Ensure you run 'make tools' first!"; exit 1; }
	@rm -fr vendor
	golangci-lint run --timeout 5m --show-stats --no-config ./...
	govulncheck --show=verbose ./...

test: fmt
	go vet ./...
	go test -count 1 -timeout 5m -v -race -cover $(TESTDIRS)

fmt-check:
	@sh -c "$(CURDIR)/scripts/gofmtcheck.sh"

fmt:
	gofmt -w $$(find . -name '*.go' |grep -v vendor)
