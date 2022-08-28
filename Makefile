SERVICE := am-wg-api
NAMESPACE := aftermath-wargaming
REGISTRY := ghcr.io/byvko-dev
# 
VERSION = "$(shell git rev-parse --short HEAD)"
TAG := ${REGISTRY}/${SERVICE}

echo:
	@echo "Tag:" ${TAG}:${VERSION}

pull:
	git pull

build:
	export DOCKER_BUILDKIT=1
	docker build -t ${TAG}:${VERSION} -t ${TAG}:latest --secret id=ssh_priv,src=$(HOME)/.ssh/id_rsa --secret id=ssh_pub,src=$(HOME)/.ssh/id_rsa.pub .
	docker image prune -f

push:
	docker push ${TAG}:latest

restart:
	kubectl rollout restart deployment/${SERVICE} -n ${NAMESPACE}
