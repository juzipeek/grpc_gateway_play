# grpc gateway 示例

## 安装包

```bash
mkdir tmp
cd tmp
# or wget https://github.com/protocolbuffers/protobuf/archive/v3.7.1.tar.gz
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install

# 下载 genproto
cd $GOPATH/src/google.golang.org
git clone https://github.com/google/go-genproto.git
mv go-genproto/ genproto/

# gopm get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

# gopm get -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

# gopm get -v github.com/golang/protobuf/protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
```

## 生成桩代码

```bash
protoc -I/usr/local/include -I. \
  -I/home/zhoucj/Project/go_play/src \
  -I/home/zhoucj/Project/go_play/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. proto/say_hi.proto
  
 protoc -I/usr/local/include -I. \
   -I/home/zhoucj/Project/go_play/src \
   -I/home/zhoucj/Project/go_play/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
   --grpc-gateway_out=logtostderr=true:. \
   proto/say_hi.proto
```

## 测试命令

```bash
curl -i -X POST -d '{"name":"jay"}' "http://localhost:8080/api/sayhi"
```