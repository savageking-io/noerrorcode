FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG AppVersion=$(cat VERSION)

RUN go build -ldflags="-X 'main.AppVersion=${AppVersion}'" -o necuser .

FROM debian:bookworm-slim AS runtime

COPY --from=builder /app/necuser /necuser
COPY --from=builder /app/VERSION /VERSION

COPY config/user.yaml /etc/noerrorcode/user.yaml

ENTRYPOINT ["/necuser", "serve", "--config", "/etc/noerrorcode/user.yaml", "--log", "trace"]

ARG AppVersion
LABEL version="${AppVersion}"
