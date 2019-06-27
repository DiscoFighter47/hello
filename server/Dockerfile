# builder image
FROM golang:latest AS builder

# copy source code
WORKDIR /go/src/github.com/DiscoFighter47/hello/server
COPY . .

# build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARC=amd64 go build -a -installsuffix cgo .


# base image
FROM alpine:latest

# Security related package
RUN apk --no-cache add ca-certificates

# copy the binary
COPY --from=builder /go/src/github.com/DiscoFighter47/hello/server/server .

# run the binary
ENTRYPOINT [ "./server" ]