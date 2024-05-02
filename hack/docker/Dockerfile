# Copyright (c) 2023 ysicing All rights reserved.
# Use of this source code is governed by WTFPL LICENSE
# license that can be found in the LICENSE file.

FROM ysicing/god AS build

WORKDIR /go/src

COPY . /go/src/

RUN task build

FROM ysicing/debian

COPY hack/docker/run.sh /run.sh

RUN chmod +x /run.sh && mkdir -p /conf /opt/go/hack

COPY example.yaml /conf/example.yaml

COPY hack/ssl /opt/go/hack/ssl

COPY --from=build /go/src/_output/bin/goexample /opt/go/goexample

WORKDIR /opt/go

ENTRYPOINT /run.sh
