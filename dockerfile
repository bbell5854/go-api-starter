# 1: Build Binary
FROM golang:1.14-alpine as build_base

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /tmp/app

COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app

# 2: Build Lightweight Container with Binary
FROM scratch

COPY --from=build_base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build_base /etc/passwd /etc/passwd
COPY --from=build_base /etc/group /etc/group
COPY --from=build_base /go/bin/app /go/bin/app

USER appuser:appuser

ENTRYPOINT ["/go/bin/app"]
