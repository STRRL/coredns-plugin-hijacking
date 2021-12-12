# syntax=docker/dockerfile:experimental

FROM golang:1.17 AS build-env
ENV GO111MODULE on
RUN git clone --depth 1 --branch v1.8.6 https://github.com/coredns/coredns
RUN echo "hijack:github.com/strrl/coredns-plugin-hijacking/hijacking" >> coredns/plugin.cfg
RUN cd coredns && make
