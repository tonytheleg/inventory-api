FROM registry.access.redhat.com/ubi9/ubi-minimal:9.6-1752587672 AS builder

ARG TARGETARCH
USER root
RUN microdnf install -y tar gzip unzip go-toolset

# Install protoc plugins
# Latest versions as of 2024-08-08
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
RUN go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@v2.0.0-20240725023016-d6fca5e3e984
RUN go install github.com/google/gnostic/cmd/protoc-gen-openapi@v0.7.0

# Install Buf
# Latest versions as of 2024-08-08
RUN go install github.com/bufbuild/buf/cmd/buf@v1.36.0

FROM registry.access.redhat.com/ubi9/ubi-minimal:9.6-1752587672

RUN microdnf install -y tar gzip unzip

COPY --from=builder /root/go/bin/protoc-gen-go /usr/local/bin/protoc-gen-go
COPY --from=builder /root/go/bin/protoc-gen-go-grpc /usr/local/bin/protoc-gen-go-grpc
COPY --from=builder /root/go/bin/protoc-gen-go-http /usr/local/bin/protoc-gen-go-http
COPY --from=builder /root/go/bin/protoc-gen-openapi /usr/local/bin/protoc-gen-openapi
COPY --from=builder /root/go/bin/buf /usr/local/bin/buf
