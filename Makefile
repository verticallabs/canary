VERSION=0.1.0
REGISTRY_HOST=pi:5000

docker-build:
	docker build -t ${REGISTRY_HOST}/verticallabs/canary:${VERSION} --build-arg GOOS=linux --build-arg GOARCH=arm --build-arg GOARM=7 .
docker-push: docker-build
	docker push ${REGISTRY_HOST}/verticallabs/canary:${VERSION}