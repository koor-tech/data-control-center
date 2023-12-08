# syntax=docker/dockerfile:1

# Go builder
FROM docker.io/library/golang:1.21-bullseye AS gobuilder
WORKDIR /app
COPY . ./
RUN apt-get update && \
    apt-get install -y git && \
    make build-go

# Ancientt
FROM docker.io/library/alpine:3.19.0 AS ancientt
ADD https://github.com/galexrt/ancientt/releases/download/v0.3.0/ancientt-0.3.0.linux-amd64.tar.gz ancientt.tar.gz
# Checksum taken from ancientt amd64 release v0.3.0 `sha256sums.txt` artifact
RUN test "84f186a2150552df4bb63892b77b6db2399464934f90cd0b3927aa3da137e0b3  ancientt.tar.gz" = "$(sha256sum ancientt.tar.gz)" && \
    tar xfvz ancientt.tar.gz --strip-components=1 && \
    rm -rf ancientt.tar.gz

# End image
FROM quay.io/ceph/ceph:v17.2.6-20230826
COPY ./.output/public ./.output/public
COPY --from=gobuilder /app/data-control-center /data-control-center
COPY --from=ancientt /ancientt /usr/local/bin/ancientt

USER nobody
COPY config.example.yaml /config/config.yaml

VOLUME /config

EXPOSE 8282/tcp

CMD ["/data-control-center"]
