FROM golang:1.20-bullseye AS builder

RUN apt update && apt install -y libcephfs-dev librbd-dev librados-dev

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  koor-data-user

WORKDIR /app

COPY . .

RUN  go mod download

RUN go build -o /api .


FROM scratch

COPY --from=builder main .

USER koor-data-user:koor-data-user

CMD ["./main"]
