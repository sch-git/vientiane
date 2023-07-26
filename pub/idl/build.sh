protoc --go_out=./grpc/ ./vientiane.proto
protoc --go-grpc_out=./grpc/ ./vientiane.proto
cd grpc
ls *.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
