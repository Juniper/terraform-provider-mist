HOSTNAME=terraform.local
NAMESPACE=local
NAME=mist
BINARY=terraform-provider-${NAME}
VERSION=0.0.1
OS_ARCH=darwin_arm64

default: install

build:
		go build -o ${BINARY}

install: build
		mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
		mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
