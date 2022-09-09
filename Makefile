SERVICE := am-wg-api
NAMESPACE := aftermath-wargaming
REGISTRY := ghcr.io/byvko-dev
# 
VERSION = $(shell git rev-parse --short HEAD)
TAG := ${REGISTRY}/${SERVICE}

echo:
	@echo "Tag:" ${TAG}:${VERSION}

pull:
	git pull

build:
	export DOCKER_BUILDKIT=1
	docker build -t ${TAG}:${VERSION} -t ${TAG}:latest --secret id=ssh_priv,src=$(HOME)/.ssh/id_rsa --secret id=ssh_pub,src=$(HOME)/.ssh/id_rsa.pub .
	docker image prune -f

build-fly:
	docker buildx build --platform linux/amd64 -t registry.fly.io/am-wg-proxy:${VERSION} -t registry.fly.io/am-wg-proxy:latest --push --secret id=ssh_priv,src=$(HOME)/.ssh/id_rsa --secret id=ssh_pub,src=$(HOME)/.ssh/id_rsa.pub .


restart:
	kubectl rollout restart deployment/${SERVICE} -n ${NAMESPACE}
	kubectl rollout restart deployment/${SERVICE}-cache -n ${NAMESPACE}
