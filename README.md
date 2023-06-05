# User service

## Run

make run.local

## Request

#### Local

##### grpcurl

grpcurl --plaintext -d '{"user": {"username": "john_doe"}}' localhost:8080 boutit.user.api.UserService/CreateUser

##### curl

curl -X POST -k http://localhost:8090/v1/user -d '{"user": {"username": "jane_doe"}}'
