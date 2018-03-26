sleeptime     ?= "5s"
resource      ?= "pod"
namespace     ?= "ansible-service-broker"
bundleid      ?= "d889087d9f39d5b09a06842518f5d9e2"
bundleparam   ?= "pod"

vendor:
	dep ensure

compile:
	go build -i -ldflags="-s -w" ./cmd/main.go

run: compile
	@SLEEPTIME=${sleeptime} RESOURCE=${resource} NAMESPACE=${namespace} BUNDLEID=${bundleid} BUNDLEPARAM=${bundleparam} ./main
