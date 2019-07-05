# builder image
FROM golang:latest AS builder

# copy source code
WORKDIR /go/src/github.com/DiscoFighter47/hello/frontend
COPY . .

# install dependencies
RUN go get -u -v github.com/golang/dep/cmd/dep
RUN dep init -v

# build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARC=amd64 go build -a -installsuffix cgo .


# base image
FROM alpine:latest

# Security related package
RUN apk --no-cache add ca-certificates

# copy the binary
COPY --from=builder /go/src/github.com/DiscoFighter47/hello/frontend/frontend .

# copy html templates
COPY ./views ./views

# run the binary
ENTRYPOINT [ "./frontend" ]