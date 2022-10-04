# User service

## Run

## Request

#### Local

##### grpcurl

grpcurl --plaintext -d '{"user": {"username": "john_doe"}}' localhost:8080 api.UserService/CreateUser
