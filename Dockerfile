FROM golang:alpine AS builder
ADD . /go/src/
WORKDIR /go/src
RUN go build -mod=vendor -ldflags "-s -w" -o ./build/bin/searching-products-wallmart-api cmd/main.go

FROM alpine
COPY --from=builder /go/src/build/bin /app/
EXPOSE 8080:8080
WORKDIR /app
CMD ["./searching-products-wallmart-api", "api"]