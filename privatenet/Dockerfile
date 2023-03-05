FROM golang:1.19-bullseye AS build-env
WORKDIR /sydney/privatenet
ADD . .
RUN go mod download
RUN make build

FROM debian:bullseye-slim
WORKDIR /sydney/privatenet
COPY --from=build-env /sydney/privatenet/.wg /sydney/privatenet/.wg
COPY --from=build-env /sydney/privatenet/bin/ /sydney/privatenet/
ENTRYPOINT [ "/bin/bash" ]
CMD [ "nohup /sydney/privatenet/ingressproxy & /sydney/privatenet/testserver" ]
