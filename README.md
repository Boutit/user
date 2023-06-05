# User service

## Run

make run.local

## Request

#### Local

##### grpcurl

grpcurl --plaintext -d '{"username": "john_doe"}' localhost:9000 boutit.user.api.UserService/SignupUser

##### curl

curl -X POST -k http://localhost:9000/v1/user -d '{"username": "jane_doe"}'
