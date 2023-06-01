FROM golang:1.19.1 AS protoc_gen_go

RUN apt update && apt install -y --no-install-recommends curl make git unzip apt-utils
ENV GO111MODULE=on
ENV PROTOC_VERSION=3.14.0
ENV GRPC_WEB_VERSION=1.2.1
ENV BUFBUILD_VERSION=0.24.0

RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOC_VERSION/protoc-$PROTOC_VERSION-linux-x86_64.zip
RUN unzip protoc-$PROTOC_VERSION-linux-x86_64.zip -d protoc3
RUN mv protoc3/bin/* /usr/local/bin/
RUN mv protoc3/include/* /usr/local/include/

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.25.0
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.0.1

# check proto syntax
RUN curl -sSL https://github.com/bufbuild/buf/releases/download/v$BUFBUILD_VERSION/buf-Linux-x86_64 -o /usr/local/bin/buf
RUN chmod +x /usr/local/bin/buf
RUN apt install -y clang-format
