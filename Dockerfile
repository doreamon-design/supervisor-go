FROM golang:alpine AS builder

RUN apk add --no-cache --update git gcc rust

COPY . /src
WORKDIR /src

RUN CGO_ENABLED=0 go build -a -ldflags "-linkmode external -extldflags -static" -o /usr/local/bin/supervisor-go github.com/ochinchina/supervisor-go

FROM scratch

COPY --from=builder /usr/local/bin/supervisor-go /usr/local/bin/supervisor-go

ENTRYPOINT ["/usr/local/bin/supervisor-go"]
