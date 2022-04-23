# Copyright (c) 2022 ysicing All rights reserved.
# Use of this source code is governed by WTFPL LICENSE
# license that can be found in the LICENSE file.

FROM registry.cn-beijing.aliyuncs.com/k7scn/god AS build

WORKDIR /go/src

ENV GOPROXY=https://goproxy.cn,direct

COPY . /go/src/

RUN make build

FROM registry.cn-beijing.aliyuncs.com/k7scn/debian

COPY hack/docker/run.sh /run.sh

RUN chmod +x /run.sh && mkdir -p /conf /opt/go/hack

COPY example.yaml /conf/example.yaml

COPY hack/ssl /opt/go/hack/ssl

COPY --from=build /go/src/_output/bin/go-example /opt/go/go-example

WORKDIR /opt/go

ENTRYPOINT /run.sh
