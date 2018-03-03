# Identifies the current build.
# These will be embedded in the app and displayed when it starts.
VERSION ?= 0.0.1.Final-SNAPSHOT
COMMIT_HASH ?= $(shell git rev-parse HEAD)

# SETTINGS BUILD
BUILD_NAME = managef_api

# Identifies the docker image that will be built and deployed.
DOCKER_ACCOUNT ?= aljesusg

DOCKER_NAME ?= ${DOCKER_ACCOUNT}/${BUILD_NAME}
DOCKER_VERSION ?= dev
DOCKER_TAG = ${DOCKER_NAME}:${DOCKER_VERSION}

# The minimum Go version that must be used to build the app.
GO_VERSION_MANAGEF = 1.8.3

# Environment variables set when running the Go compiler.
GO_BUILD_ENVVARS = \
	GOOS=linux \
	GOARCH=amd64 \
    CGO_ENABLED=0 \

all: build

clean:
	@echo Cleaning...
	@rm -f sws
	@rm -rf ${GOPATH}/bin/${BUILD_NAME}
	@rm -rf ${GOPATH}/pkg/*
	@rm -rf _output/*

go-check:
	@hack/check_go_version.sh "${GO_VERSION_MANAGEF}"

build: go-check
	@echo Building...
	${GO_BUILD_ENVVARS} go build \
		-o ${GOPATH}/bin/${BUILD_NAME} -ldflags "-X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH}"

install:
	@echo Installing...
	${GO_BUILD_ENVVARS} go install \
		-ldflags "-X main.version=${VERSION} -X main.commitHash=${COMMIT_HASH}"


.prepare-docker-image-files:
	@echo Preparing docker image files...
	@mkdir -p ${GOPATH}/_output/docker/${BUILD_NAME}
	@cp -r deploy/docker/* ${GOPATH}/_output/docker/${BUILD_NAME}/
	@cp ${GOPATH}/bin/${BUILD_NAME} ${GOPATH}/_output/docker/${BUILD_NAME}/
	@echo ${DOCKER_TAG}

docker: .prepare-docker-image-files
	@echo Building docker image into local docker daemon...
	docker build -t ${DOCKER_TAG} ${GOPATH}/_output/docker/${BUILD_NAME}

docker-push:
	@echo Pushing current docker image to ${DOCKER_TAG}
	docker push ${DOCKER_TAG}