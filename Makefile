generate-pb:
	@docker run -v `pwd`:/defs namely/protoc-all -l go -f proto/service.proto --with-gateway
