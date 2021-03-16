BINARY = "app"
BUILD_SHA=`git rev-parse HEAD`
BUILD_TIME=`date +%Y%m%d%H%M%S`
LDFLAGS=-ldflags "-X main.build_id=${BUILD_SHA} -X main.build_time=${BUILD_TIME}"

build-local:
	go build ${LDFLAGS} -o ${BINARY}

build-darwin:
	go build ${LDFLAGS} -o darwin-${BINARY}

build-linux: $(GOFILES)
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o linux-${BINARY}

build-local-docker-image: build-linux
	docker build -t chris/circleci-golang-demo:coolest .

run-image:
	# container:host
	docker run -p 81:8080 chris/circleci-golang-demo:coolest

reset:
	echo "Reseting demo!"
