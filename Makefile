_ := $(shell mkdir -p .make bin)

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

DEVOPS   := ${LOCALBIN}/devops
GOLANGCI := ${LOCALBIN}/golangci-lint

export GOBIN := ${LOCALBIN}

build: bin/thecluster
tidy: go.sum
lint: .make/lint
format: .make/format
init: bin/devops bin/golangci-lint

go.sum: go.mod $(shell $(DEVOPS) list --go) | bin/devops
	go mod tidy

bin/thecluster: go.mod $(shell $(DEVOPS) list --go) | bin/devops
	go -C cmd/thecluster build -o ${WORKING_DIR}/$@

bin/devops: .versions/devops
	go install github.com/unmango/go/cmd/devops@v$(shell cat $<)

bin/golangci-lint: .versions/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${LOCALBIN} v$(shell cat $<)

.envrc: hack/example.envrc
	cp $< $@

.make/lint: $(shell $(DEVOPS) list --go) | bin/golangci-lint
	$(GOLANGCI) run $(sort $(dir $?))
	@touch $@

.make/format: $(shell $(DEVOPS) list --go --absolute)
	go fmt $(sort $(dir $?))
	@touch $@
