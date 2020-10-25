FROM registry.cn-beijing.aliyuncs.com/k7scn/dgo AS build

WORKDIR /go/src

ENV GOPROXY=https://goproxy.cn,direct

RUN apt update && apt install build-essential -y

COPY . /go/src/

RUN go get github.com/mitchellh/gox && make build

FROM registry.cn-beijing.aliyuncs.com/k7scn/debian

COPY --from=build /go/src/dist/go-example /root/go-example

COPY static /root/static

RUN chmod +x /root/go-example

WORKDIR /root

ENTRYPOINT /root/go-example