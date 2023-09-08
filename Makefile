selfVer=1.0
commit=$(shell git rev-parse --short HEAD)
date=$(shell date +'%Y%m%d-%H:%M:%S')
gitversion=$(selfVer)-$(commit)-$(date)
goversion=$(shell go version | awk '{print $$3}')
version=$(gitversion)-$(goversion)

NAMESPACE=mds
COMMIT=$(shell git rev-parse --short HEAD)
DATE:=$(shell date +'%Y%m%d-%H%M%S')
QCI_ENV_FILE?=release-tag
IMG-DataDog=data_dog
OBJ-DataDog=DataDog
BIN-DataDog=bin

TAG ?= release-init-$(DATE)-$(COMMIT)

build:
	GOARCH=amd64 GOOS=linux go build -ldflags "-X main.Version=$(version)" -mod=vendor -o $(BIN-DataDog)/$(OBJ-DataDog) main.go

build-image:build
	mv $(BIN-DataDog)/$(OBJ-DataDog) docker/
	cd docker; docker build --rm --no-cache -t csighub.tencentyun.com/$(NAMESPACE)/$(IMG-DataDog):$(TAG) .

push-image:build-image
	docker push csighub.tencentyun.com/$(NAMESPACE)/$(IMG-DataDog):$(TAG)
	echo "TKEx_Image_Hub=csighub.tencentyun.com/$(NAMESPACE)/$(IMG-DataDog):$(TAG)" >> ${QCI_ENV_FILE}
