SERVER_PATH="cmd/server/main.go"
PROTO_RELATIVE_PATH="api/v1"

codegen.proto:
	protoc $(PROTO_RELATIVE_PATH)/*.proto --go_out=. --go_opt=paths=source_relative --proto_path=.

run:
	go run $(SERVER_PATH)