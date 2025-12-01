YEAR ?= $(shell date "+%Y")
DAY ?= $(shell date "+%d")
DAY_DIR ?= ./${YEAR}/${DAY}

.ONESHELL:

new:
	mkdir -p ${DAY_DIR}
	cd ${DAY_DIR} && \
		go mod init "github.com/tsatam/adventofcode/${YEAR}/${DAY}" && \
		cd ../..
	go work use ${DAY_DIR}

test:
	go test "${DAY_DIR}/..."

run:
	go run ${DAY_DIR}