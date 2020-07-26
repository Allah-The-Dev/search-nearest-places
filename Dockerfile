FROM golang:1.14 as builder
WORKDIR /go/src/app
COPY . .
RUN go get -v ./...
RUN go build -o app

FROM alpine:latest
WORKDIR /go/src/app
COPY --from=builder /go/src/app/app .
CMD ["/.app"]
EXPOSE 9080
