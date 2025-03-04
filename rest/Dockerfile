FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ARG AppVersion=$(cat VERSION)

RUN go build -ldflags="-X 'main.AppVersion=${AppVersion}'" -o app .

FROM gcr.io/distroless/base-debian11

COPY --from=builder /app/app /app
COPY --from=builder /app/VERSION /VERSION

COPY config/rest.yaml /etc/noerrorcode/rest.yaml

ENTRYPOINT ["/app", "serve", "--config", "/etc/noerrorcode/rest.yaml"]

ARG AppVersion
LABEL version="${AppVersion}"
