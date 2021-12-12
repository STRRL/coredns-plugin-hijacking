# syntax=docker/dockerfile:experimental

FROM golang:1.17 AS build-env
ENV GO111MODULE on
WORKDIR /
RUN git clone --depth 1 --branch v1.8.6 https://github.com/coredns/coredns
RUN echo "hijack:github.com/strrl/coredns-plugin-hijacking/hijacking" >> coredns/plugin.cfg
RUN cd coredns && make

FROM debian:stable-slim AS certs
RUN apt-get update && apt-get -uy upgrade
RUN apt-get -y install ca-certificates && update-ca-certificates

FROM scratch
LABEL org.opencontainers.image.source=https://github.com/STRRL/coredns-plugin-hijacking
COPY --from=certs /etc/ssl/certs /etc/ssl/certs
COPY --from=build-env /coredns/coredns /coredns
EXPOSE 53 53/udp
ENTRYPOINT ["/coredns"]
