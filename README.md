# User service

## Run

make run.local

## Request

#### Local

##### grpcurl

grpcurl --plaintext -d '{"username": "john_doe"}' localhost:9000 boutit.user.api.UserService/SignupUser

##### curl

Note: curl must be done through api-gateway

curl -X POST -k http://localhost:8090/v1/signup -d '{"email_or_phone": "6472289484"}'
