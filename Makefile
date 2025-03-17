_ := $(shell mkdir -p .make bin)

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

DEVCTL   := go tool devctl
GINKGO   := go tool ginkgo
GOLANGCI := ${LOCALBIN}/golangci-lint

export GOBIN := ${LOCALBIN}

GO_SRC := $(shell $(DEVCTL) list --go)

ifeq ($(CI),)
TEST_FLAGS := --label-filter '!E2E'
else
TEST_FLAGS := --github-output --race --trace
endif

build: bin/thecluster
test: .make/test
tidy: go.sum
lint: .make/lint
format: .make/format
init: bin/golangci-lint

test_all:
	$(GINKGO) run -r ./

%_suite_test.go: | bin/ginkgo
	cd $(dir $@) && $(GINKGO) bootstrap

$(GO_SRC:%.go=%_test.go): %_test.go:
	cd $(dir $@) && $(GINKGO) generate $(notdir $*)

go.sum: go.mod ${GO_SRC}
	go mod tidy

bin/thecluster: go.mod ${GO_SRC}
	go build -o ${WORKING_DIR}/$@

bin/golangci-lint: .versions/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${LOCALBIN} v$(shell cat $<)

.envrc: hack/example.envrc
	cp $< $@

.make/lint: ${GO_SRC} | bin/golangci-lint
	$(GOLANGCI) run $(sort $(dir $?))
	@touch $@

.make/format: $(shell $(DEVCTL) list --go --absolute)
	go fmt $(sort $(dir $?))
	@touch $@

.make/test: ${GO_SRC}
	$(GINKGO) run ${TEST_FLAGS} $(sort $(dir $?))
	@touch $@
