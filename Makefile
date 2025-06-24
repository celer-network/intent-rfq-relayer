install:
	@echo "--> Ensure dependencies have not been modified"
	GO111MODULE=on go mod verify
	go build -o $(HOME)/go/bin/rfq-relayer ./relayer/main/main.go

.PHONY: proto
proto:
	protoc --proto_path=./relayer/proto --go_out ./relayer --go_opt=module=github.com/celer-network/intent-rfq-relayer/relayer \
	--go-grpc_out=./relayer --go-grpc_opt=require_unimplemented_servers=false,module=github.com/celer-network/intent-rfq-relayer/relayer \
	--grpc-gateway_out=./relayer --grpc-gateway_opt=module=github.com/celer-network/intent-rfq-relayer/relayer \
	--openapiv2_out ./relayer/openapi \
	./relayer/proto/service/*/*.proto