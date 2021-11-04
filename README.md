# grpc + grpc gateway demo

## Preparation
11
OS: Windows 10

1. 安装 `Protocol buffer compiler`，下载地址：[protobuf releases](https://github.com/protocolbuffers/protobuf/releases)，并设置环境变量。

2. 安装 `protocol compiler` Go 插件，用于生成 `*.pb.go` 文件：

    ```shell script
    go get -u github.com/golang/protobuf/protoc-gen-go
    ```

    并将`GOPATH/bin`加入环境变量中。

3. 安装 `grpc-gateway` ，用于生成 `*.pb.gw.go` 和 `*.swagger.json` 文件

    ```shell script
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
    ```

## Run

### 编译 proto 文件

-IE:\go\include -IE:\go\src
生成 grpc 相关文件 `.pb.go`，其中包括 用于序列化我们请求响应数据的代码、grpc 客户端 和 grpc 服务端：

```shell script
protoc -I%PROTOC_INCLUDE% -I%GOPATH%\src -Ithird_party\googleapis -I. --go_out=plugins=grpc:. api\helloworld.proto
```

生成 gateway 文件 `.pb.gw.go`：

```shell script
protoc -I%PROTOC_INCLUDE% -I%GOPATH%\src -Ithird_party\googleapis -I. --grpc-gateway_out=logtostderr=true:. api\helloworld.proto
```

生成 swagger 文档 `.swagger.json`：

```shell script
protoc -I%PROTOC_INCLUDE% -I%GOPATH%\src -Ithird_party\googleapis -I. --swagger_out=logtostderr=true:. api\helloworld.proto
```

### 运行

运行 grpc 服务端：

```shell script
go run server\server.go
```

打开另一个命令行，运行 grpc 客户端：

```shell script
go run client\client.go
```

再打开另一个命令行，运行 HTTP 网关：

```shell script
go run gateway\gateway.go
```

## Run with docker

### 生成镜像

```shell script
docker build --target server -t <server_image_name> .
docker build --target gateway -t <gateway_image_name> .
```

### 构建并运行容器

本地运行时，可以创建容器网络，并将两个容器接入网络，则相互之间可通过**容器名**访问。

```shell script
docker network create my-net
docker run -d --name server --network my-net -p 8081:8081 <server_image_name>
docker run -d --name gateway --network my-net -p 8080:8080 <gateway_image_name> /root/Gateway --server_addr server:8081
```

分布式运行容器，则需要提供 `server` 的 IP 地址：

```shell script
docker run -d --name server -p 8081:8081 <server_image_name>
docker run -d --name gateway -p 8080:8080 <gateway_image_name> /root/Gateway --server_addr <server_ip>:8081
```
