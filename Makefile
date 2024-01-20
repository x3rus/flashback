
DOCKER_PREFIX=x3rus/
APP_NAME=flashback
APP_VERSION=0.1
LOCAL_ALBUMS="${PWD}/data/pic-sample/"
MOUNTED_ALBUMS="/data/pic-sample/"


build: 
	go build -o ${APP_NAME} src/*.go 

docker-build: 
	docker build -t ${DOCKER_PREFIX}${APP_NAME}:${APP_VERSION} .


docker-run: docker-build
	docker run --rm -e ALBUMDIRS=${MOUNTED_ALBUMS} -v ${LOCAL_ALBUMS}:${MOUNTED_ALBUMS} \
		-p 8085:8080 ${DOCKER_PREFIX}${APP_NAME}:${APP_VERSION}

bench:
	go test -bench=. -count=2 ./src/
	go test -v -tags=benchmark ./src/

test:
	go test -v -tags=default ./src/

alltest:
	go test -v -tags=all ./src/

setup-linter:
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.2

lint: setup-linter
	go vet  ./src/
	$(go env GOPATH)/bin/golangci-lint run ./src/

localdev: build
	ALBUMDIRS="data/pic-sample/" ./flashback
