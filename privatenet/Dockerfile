FROM golang:1.19-bullseye AS build-env
WORKDIR /sydney/privatenet
ADD . .
RUN go mod download
RUN make build

FROM debian:bullseye-slim
WORKDIR /sydney/privatenet
RUN apt-get update \
  && apt-get install -y iperf3 \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*
COPY --from=build-env /sydney/privatenet/.wg /sydney/privatenet/.wg
COPY --from=build-env /sydney/privatenet/bin/ /sydney/privatenet/
COPY --from=build-env /sydney/privatenet/init.sh /sydney/privatenet/init.sh
ENTRYPOINT /bin/sh
CMD init.sh
