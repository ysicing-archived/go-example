FROM registry.cn-beijing.aliyuncs.com/k7scn/dgo AS build

WORKDIR /go/src

ENV GOPROXY=https://goproxy.cn,direct

RUN apt update && apt install build-essential -y

COPY . /go/src/

RUN go get github.com/mitchellh/gox && make build

FROM registry.cn-beijing.aliyuncs.com/k7scn/debian

COPY hack/docker/server/run.sh /run.sh

RUN chmod +x /run.sh && mkdir -p /conf /opt/go

COPY --from=build /go/src/go-example /opt/go/example

WORKDIR /opt/go

ENTRYPOINT /run.sh