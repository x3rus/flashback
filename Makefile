
APP_NAME=flashback
APP_VERSION=0.1
LOCAL_ALBUMS="${PWD}/data/pic-sample/"
MOUNTED_ALBUMS="/data/pic-sample/"


build: 
	go build -o ${APP_NAME} src/*.go 

docker-build: 
	docker build -t ${APP_NAME}:${APP_VERSION} .


docker-run: docker-build
	docker run --rm -e ALBUMDIRS=${MOUNTED_ALBUMS} -v ${LOCAL_ALBUMS}:${MOUNTED_ALBUMS} \
		            -p 8080 ${APP_NAME}:${APP_VERSION}

test:
	go test -v src/*.go

lint:
	go vet  src/*.go

localdev: build
	ALBUMDIRS="data/pic-sample/" ./flashback
