_ := $(shell mkdir -p .make bin)

WORKING_DIR := $(shell pwd)
LOCALBIN    := ${WORKING_DIR}/bin

DEVOPS := ${LOCALBIN}/devops

export GOBIN := ${LOCALBIN}

build: bin/thecluster
tidy: go.sum

go.sum: go.mod $(shell $(DEVOPS) list --go) | bin/devops
	go mod tidy

bin/thecluster: go.mod $(shell $(DEVOPS) list --go) | bin/devops
	go -C cmd/thecluster build -o ${WORKING_DIR}/$@

bin/devops: .versions/devops
	go install github.com/unmango/go/cmd/devops@v$(shell cat $<)

.envrc: hack/example.envrc
	cp $< $@
