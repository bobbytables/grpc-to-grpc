generate-pb:
	@docker run -v `pwd`:/defs namely/protoc-all -l go -f proto/service.proto --with-gateway

docker.gw:
	@docker build -f Dockerfile.gw -t bobbytables/jokes-gw:latest .

docker.grpc:
	@docker build -f Dockerfile.grpc -t bobbytables/jokes-grpc:latest .

docker.envoy:
	@docker build -f Dockerfile.envoy -t bobbytables/envoy-minikube-test:latest .

docker.init:
	@docker build -f Dockerfile.init -t bobbytables/envoy-minikube-init:latest .
