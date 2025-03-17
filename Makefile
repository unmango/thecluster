_ := $(shell mkdir -p .make bin)

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

DEVOPS   := ${LOCALBIN}/devops
GINKGO   := ${LOCALBIN}/ginkgo
GOLANGCI := ${LOCALBIN}/golangci-lint

export GOBIN := ${LOCALBIN}

GO_SRC := $(shell $(DEVOPS) list --go)

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
init: bin/devops bin/golangci-lint

test_all:
	$(GINKGO) run -r ./

%_suite_test.go: | bin/ginkgo
	cd $(dir $@) && $(GINKGO) bootstrap

$(GO_SRC:%.go=%_test.go): %_test.go: | bin/ginkgo
	cd $(dir $@) && $(GINKGO) generate $(notdir $*)

go.sum: go.mod ${GO_SRC} | bin/devops
	go mod tidy

bin/thecluster: go.mod ${GO_SRC} | bin/devops
	go -C cmd build -o ${WORKING_DIR}/$@

bin/devops: .versions/devops
	go install github.com/unmango/go/cmd/devops@v$(shell cat $<)

bin/golangci-lint: .versions/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${LOCALBIN} v$(shell cat $<)

bin/ginkgo: go.mod
	go install github.com/onsi/ginkgo/v2/ginkgo

.envrc: hack/example.envrc
	cp $< $@

.make/lint: ${GO_SRC} | bin/golangci-lint
	$(GOLANGCI) run $(sort $(dir $?))
	@touch $@

.make/format: $(shell $(DEVOPS) list --go --absolute)
	go fmt $(sort $(dir $?))
	@touch $@

.make/test: ${GO_SRC} | bin/ginkgo
	$(GINKGO) run ${TEST_FLAGS} $(sort $(dir $?))
	@touch $@
