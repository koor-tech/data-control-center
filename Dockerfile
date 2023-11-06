# syntax=docker/dockerfile:1
FROM golang:1.21-bullseye AS gobuilder
WORKDIR /app
COPY . ./
RUN apt-get update && \
    apt-get install -y git && \
    make build-go

FROM quay.io/ceph/ceph:v17.2.6-20230826
COPY ./.output/public ./.output/public
COPY --from=gobuilder /app/data-control-center /data-control-center
USER nobody
COPY config.example.yaml /config/config.yaml
VOLUME /config

EXPOSE 8282/tcp

CMD ["/data-control-center"]
