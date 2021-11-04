FROM golang:1.14 as builder

MAINTAINER dounine

WORKDIR /go/grpcDemo

COPY . /go/grpcDemo

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn \
    && echo export PATH="$PATH:$(go env GOPATH)/bin" >> ~/.bashrc

RUN apt-get update\
    && apt install -y protobuf-compiler \
    && go get -u github.com/golang/protobuf/protoc-gen-go \
    && go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway

RUN protoc -I. -I third_party/googleapis/ --go_out=plugins=grpc:. api/helloworld.proto \
    && protoc -I. -I third_party/googleapis/ --grpc-gateway_out=logtostderr=true:. api/helloworld.proto

RUN go build -o Server server/server.go \
    && go build -o Gateway gateway/gateway.go


FROM centos as server

WORKDIR /root

COPY --from=builder /go/grpcDemo/Server .

CMD ["./Server"]

EXPOSE 9061


FROM centos as gateway

WORKDIR /root

COPY --from=builder /go/grpcDemo/Gateway .

CMD ["./Gateway", "--server_addr", "server:9061"]

EXPOSE 9060