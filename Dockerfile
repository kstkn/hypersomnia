FROM golang:1.12 AS builder

RUN mkdir /app

COPY . /app/

WORKDIR /app

RUN env

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM scratch

COPY --from=builder /app/hypersomnia .

CMD ["/hypersomnia"]
