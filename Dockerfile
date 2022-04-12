FROM golang:alpine as gobuild
RUN apk add git
WORKDIR /src
COPY go.* ./
RUN go mod download
COPY *.go ./
RUN go get golang.org/x/crypto/acme/autocert
RUN go build ./

FROM alpine:latest
COPY --from=gobuild /src /
WORKDIR /
EXPOSE 8080
ENTRYPOINT ["/go-hello"]