# golang-grpc-bcrypt-service
## Generate Proto files
```bash
protoc -I golangbcryptservice/ golangbcryptservice/bcrypt.proto --go_out=plugins=grpc:golangbcryptservice
```
