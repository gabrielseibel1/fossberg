all: protocol verify

.PHONY: protocol
protocol: 
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protocol/fossberg.proto

verify: vet check test

.PHONY: vet
vet:
	go vet ./...

.PHONY: check
check:
	staticcheck ./...

.PHONY: test
test:
	go test --count 1 --cover --coverprofile=./cover.out ./...