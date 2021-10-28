from golang:1.16 AS builder
RUN mkdir -p /go/src/github.com/x3rus/flashback/
WORKDIR /go/src/github.com/x3rus/flashback/
RUN apt-get update && \
    apt-get install -y libexif-dev && \
    apt-get clean all
COPY go.mod /go/src/github.com/x3rus/flashback/
COPY src /go/src/github.com/x3rus/flashback/src
RUN go get -v ...
RUN CGO_ENABLED=0 GOOS=linux go build  -o flashback src/*.go 

FROM alpine:latest
RUN apk --no-cache add ca-certificates exiftool
WORKDIR /root/
COPY --from=builder /go/src/github.com/x3rus/flashback/flashback ./
EXPOSE 8080
ENTRYPOINT ["./flashback"]
