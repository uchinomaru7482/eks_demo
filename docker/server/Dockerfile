FROM golang:1.19

WORKDIR /app

COPY ./server /app

RUN go build -ldflags="-w -s" -o /go/bin/app cmd/main.go

ENTRYPOINT ["/go/bin/app"]
