APP_NAME := hkg-collector
BUILD_VERSION   := $(shell git tag --contains)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD )

.PHONY: build
build:
	go build -ldflags \
		"\
		-X 'main.BuildVersion=${BUILD_VERSION}' \
		-X 'main.BuildTime=${BUILD_TIME}' \
		-X 'main.CommitID=${COMMIT_SHA1}' \
		"\
		-o ./bin/${APP_NAME}

.PHONY: run
run:
	./bin/${APP_NAME}

.PHONY: run-fs
run-fs:
	MSA_CONFIG_DEFINE='{"source":"file","prefix":"/etc/msa/","key":"collector.yml"}' ./bin/${APP_NAME}

.PHONY: run-cs
run-cs:
	MSA_CONFIG_DEFINE='{"source":"consul","prefix":"/xtc/hkg/config","key":"collector.yml"}' ./bin/${APP_NAME}

.PHONY: call
call:
	MICRO_REGISTRY=consul micro call xtc.api.hkg.collector Healthy.Echo '{"msg":"hello"}'
	MICRO_REGISTRY=consul micro call xtc.api.hkg.collector Document.Scrape '{"name":"w3c", "keyword":["web", "www"], "address":"https://baike.baidu.com/item/w3c", "attribute":"div.main-content"}'
	MICRO_REGISTRY=consul micro call xtc.api.hkg.collector Document.List 

.PHONY: post
post:
	curl -X POST -d '{"msg":"hello"}' localhost:8080/hkg/collector/Healthy/Echo

.PHONY: bm
bm:
	python3 benchmark.py

.PHONY: dist
dist:
	mkdir dist
	tar -zcf dist/${APP_NAME}-${BUILD_VERSION}.tar.gz ./bin/${APP_NAME}
