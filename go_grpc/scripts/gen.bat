echo "gen rpc server code"

SET OUT="..\server\rpc"
protoc ^
--go_out=%OUT% ^
--go-grpc_out=%OUT% ^
--go-grpc_opt=require_unimplemented_servers=false ^
server.proto


echo "gen rpc client code"

SET OUT="..\client\rpc"
protoc ^
--go_out=%OUT% ^
--go-grpc_out=%OUT% ^
--go-grpc_opt=require_unimplemented_servers=false ^
server.proto