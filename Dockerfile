FROM golang:1.14 AS builder
WORKDIR /go/src/seva
EXPOSE 8080

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -a ./cmd/seva

FROM alpine:3.11.6
WORKDIR /root/
COPY --from=builder /go/src/seva/seva .
CMD ["./seva"]
