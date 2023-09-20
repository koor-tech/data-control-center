# syntax=docker/dockerfile:1
FROM docker.io/library/node:20.7.0-alpine3.17 AS nodebuilder
WORKDIR /app
COPY . ./
RUN rm -rf ./.nuxt/ && \
    apk add --no-cache git && \
    yarn && yarn generate

FROM golang:1.21-bullseye AS gobuilder
RUN apt update && \
  apt install -y lsb-release && \
  wget -q -O- 'https://download.ceph.com/keys/release.asc' | apt-key add - && \
  echo deb https://download.ceph.com/debian-quincy/ $(lsb_release -sc) main | tee /etc/apt/sources.list.d/ceph.list && \
  apt update && \
  apt install -y libcephfs-dev librbd-dev librados-dev
WORKDIR /app
COPY . ./

RUN  go mod download && \
  go build -o /data-control-center .

FROM quay.io/ceph/ceph:v17.2.6-20230826
COPY --from=nodebuilder /app/.output/public ./.output/public
COPY --from=gobuilder /data-control-center /data-control-center
USER nobody
COPY config.example.yaml /config/config.yaml
VOLUME /config

EXPOSE 8282/tcp

CMD ["/data-control-center"]
