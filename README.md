# Install 
```sh
brew install protoc-gen-go
brew install protoc-gen-go-grpc
```

# Usage
(1) Update the proto file in /pbfiles folder

(2) Under the same /pbfiles, Run the command and auto generate Go code from your Protocol Buffer files.
```sh
protoc --proto_path=. --go_out=../ai_service --go_opt=paths=source_relative --go-grpc_out=../ai_service --go-grpc_opt=paths=source_relative ai.proto
protoc --proto_path=. --go_out=../s3_service --go_opt=paths=source_relative --go-grpc_out=../s3_service --go-grpc_opt=paths=source_relative s3.proto
```
(3) Under the GO file folder, remove the go.mod and go.sum and re-generate the dependencies modules file
```sh
go mod init github.com/poy-web3/protos/ai_service
go mod init github.com/poy-web3/protos/s3_service
go mod tidy
```
(4) Include all change file in your commit

(5) After merge the change, update the modeul in upstream service
```sh
go update github.com/poy-web3/protos/notification_service
```
or
```sh
go get github.com/poy-web3/protos/notification_service
```
