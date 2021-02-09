BINARY = "app"
CIRCLE_SHA1="123"
BUILD_TIME=`date +%Y%m%d%H%M%S`
LDFLAGS=-ldflags "-X main.build_id=${CIRCLE_SHA1} -X main.build_time=${BUILD_TIME}"

build-local:
	go build ${LDFLAGS} -o ${BINARY}

reset:
	echo "Reseting demo!"