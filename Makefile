APP_NAME := am-wg-api
REGISTRY := registry.fly.io
# 
VERSION = $(shell git rev-parse --short HEAD)
TAG := ${REGISTRY}/${APP_NAME}

echo:
	@echo "Tag:" ${TAG}:${VERSION}

build:
	docker buildx build --platform linux/amd64 -t ${TAG}:${VERSION} -t ${TAG}:latest --push --secret id=ssh_priv,src=$(HOME)/.ssh/id_rsa --secret id=ssh_pub,src=$(HOME)/.ssh/id_rsa.pub .

deploy:
	flyctl deploy --app ${APP_NAME} --image ${TAG}:${VERSION}