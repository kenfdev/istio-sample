FROM golang:1.9 as builder
WORKDIR /go/src/github.com/kenfdev/istio-sample/api/api1

COPY vendor vendor
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/kenfdev/istio-sample/api/api1/app .
CMD ["./app"]