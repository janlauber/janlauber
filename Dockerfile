# syntax=docker/dockerfile:1
FROM golang:1.17
WORKDIR /go/src/github.com/janlauber/janlauber/
COPY . .
RUN go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/janlauber/janlauber/app ./
EXPOSE 8080
CMD ["./app"]