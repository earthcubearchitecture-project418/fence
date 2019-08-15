BINARY := fence
DOCKERVER :=`cat VERSION`
.DEFAULT_GOAL := linux

linux:
		cd cmd/$(BINARY) ; \
		GOOS=linux GOARCH=amd64 CGO_ENABLED=0 env go build -o $(BINARY)

docker:
		docker build  --tag="fils/p418fence:$(DOCKERVER)"  --file=./build/Dockerfile .

dockerlatest:
		docker build  --tag="fils/p418fence:latest"  --file=./build/Dockerfile .

publish:  
		docker push fils/p418fence:$(DOCKERVER)

tag:
	docker tag fils/p418fence:$(DOCKERVER) gcr.io/top-operand-112611/p418fence:$(DOCKERVER)

publishgcr:
	docker push gcr.io/top-operand-112611/p418fence:$(DOCKERVER)

togcr: linux docker tag publishgcr