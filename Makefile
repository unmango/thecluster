_ := $(shell mkdir -p .make bin)

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

DEVCTL    := go tool devctl
GINKGO    := go tool ginkgo
GOLANGCI  := ${LOCALBIN}/golangci-lint
WATCHEXEC := ${LOCALBIN}/watchexec

export GOBIN := ${LOCALBIN}

GO_SRC != $(DEVCTL) list --go

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
init: build bin/golangci-lint

start: bin/thecluster
	$<

test_all:
	$(GINKGO) run -r ./

golden:
	$(GINKGO) run -r ./app -- -update

# https://github.com/charmbracelet/bubbletea/issues/150#issuecomment-2492038498
watch: | bin/watchexec
	$(WATCHEXEC) -e go -r --clear --wrap-process session -- 'go run .'

%_suite_test.go:
	cd $(dir $@) && $(GINKGO) bootstrap

$(GO_SRC:%.go=%_test.go): %_test.go:
	cd $(dir $@) && $(GINKGO) generate $(notdir $*)

go.sum: go.mod ${GO_SRC}
	go mod tidy

bin/thecluster: go.mod ${GO_SRC}
	go build -o ${WORKING_DIR}/$@

bin/watchexec: | .make/watchexec/watchexec
	ln -s ${CURDIR}/$| ${CURDIR}/$@

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

.make/watchexec/watchexec: .make/watchexec.tar.xz
	mkdir -p $(dir $@) && tar -C $(dir $@) -xvf $< --strip-components=1
	@touch $@

.make/watchexec.tar.xz: .versions/watchexec
ifeq ($(shell go env GOOS),darwin)
	curl -Lo $@ https://github.com/watchexec/watchexec/releases/download/$(shell $(DEVCTL) $<)/watchexec-$(shell $(DEVCTL) v watchexec)-aarch64-apple-darwin.tar.xz
else
	curl -Lo $@ https://github.com/watchexec/watchexec/releases/download/$(shell $(DEVCTL) $<)/watchexec-$(shell $(DEVCTL) v watchexec)-x86_64-unknown-linux-gnu.tar.xz
endif
