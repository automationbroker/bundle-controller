sleeptime     ?= "5s"
resource      ?= "pod"
namespace     ?= "ansible-service-broker"
bundlename    ?= "dh-dynamic-apb"
bundleid      ?= "d889087d9f39d5b09a06842518f5d9e2"
bundleparam   ?= "pods"

vendor:
	dep ensure

compile:
	go build -i -ldflags="-s -w" ./cmd/main.go

build: compile
	docker build --tag ansibleplaybookbundle/bundle-controller:latest .

run: compile
	@SLEEPTIME=${sleeptime} RESOURCE=${resource} NAMESPACE=${namespace} BUNDLEID=${bundleid} BUNDLENAME=${bundlename} BUNDLEPARAM=${bundleparam} ./main

.PHONY: run vendor compile build
