#Install 
```sh
brew install protoc-gen-go
brew install protoc-gen-go-grpc
```

#Usage
(1) Update the proto file in /pbfiles
(2) Under the /pbfiles, Run the command and auto generate Go code from your Protocol Buffer files.
```sh
protoc --proto_path=. --go_out=../notification_service --go_opt=paths=source_relative --go-grpc_out=../notification_service --go-grpc_opt=paths=source_relative notification.proto
```
