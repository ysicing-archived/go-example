BUILD_VERSION   ?= $(shell cat version.txt || echo "0.0.1")
BUILD_DATE      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse --short HEAD || echo "0.0.0")

help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

fmt:

	@echo gofmt -l
	@OUTPUT=`gofmt -l . 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "gofmt must be run on the following files:"; \
        echo "$$OUTPUT"; \
        exit 1; \
    fi

lint:

	@echo golangci-lint run ./...
	@OUTPUT=`command -v golangci-lint >/dev/null 2>&1 && golangci-lint run ./... 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "golangci-lint errors:"; \
		echo "$$OUTPUT"; \
		exit 1; \
	fi

default: fmt lint ## fmt code

static: ## 构建ui
	hack/build/genui.sh

ui: static ## 将ui编译为go文件
	hack/build/genui2go.sh

doc: ## doc
	hack/build/gendocs.sh

build: ## 构建二进制
	@echo "build bin ${BUILD_VERSION} ${BUILD_DATE} ${COMMIT_SHA1}"
	#@bash hack/docker/build.sh ${version} ${tagversion} ${commit_sha1}
	# go get github.com/mitchellh/gox
	@CGO_ENABLED=1 GOARCH=amd64 go build -o dist/go-example \
    	-ldflags   "-X 'app/constants.Commit=${COMMIT_SHA1}' \
                    -X 'app/constants.Date=${BUILD_DATE}' \
                    -X 'app/constants.Release=${BUILD_VERSION}'"

docker: ## 构建镜像
	# hack/build/gendocs.sh
	# hack/build/genui.sh
	# hack/build/genui2go.sh
	docker build -t goexample:${BUILD_VERSION} -f hack/docker/server/Dockerfile  .

clean: ## clean
	rm -rf dist/*
	rm -rf ui/dist

.PHONY : build release clean

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://goproxy.cn
GOSUMDB = sum.golang.google.cn