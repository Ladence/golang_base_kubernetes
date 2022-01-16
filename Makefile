GOOS?=linux
GOARCH?=amd64

APP?=go-kube-service
PORT?=8080
PROJECT?=github.com/Ladence/golang_base_kubernetes
DOCKER_IMAGE?=docker.io/ladence/${APP}

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

clean:
	rm -f ${APP}

build: clean
	GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/internal/version.Release=${RELEASE} \
		-X ${PROJECT}/internal/version.Commit=${COMMIT} -X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

container: build
	docker build -t ${DOCKER_IMAGE}:${RELEASE} .

push: container
	docker push ${DOCKER_IMAGE}:${RELEASE}

run: container
	docker stop ${APP}:${RELEASE} || true && docker rm ${APP}:${RELEASE} || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		   -e "PORT=${PORT}" ${APP}:${RELEASE}

minikube: push
	for t in $(shell find ./kube/ -type f -name "*.yaml"); do \
        cat $$t | \
        	gsed -E "s/\{\{(\s*)\.Release(\s*)\}\}/$(RELEASE)/g" | \
        	gsed -E "s/\{\{(\s*)\.ServiceName(\s*)\}\}/$(APP)/g"; \
        echo ; \
    done > tmp.yaml
	kubectl apply -f tmp.yaml

test:
	go test -v -race ./...