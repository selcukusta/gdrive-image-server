FROM golang:1.14-alpine as builder
COPY . $GOPATH/src/github.com/selcukusta/simple-image-server
WORKDIR $GOPATH/src/github.com/selcukusta/simple-image-server/cmd/image-server
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o $GOPATH/bin/simple-image-server .
FROM scratch as final
COPY --from=builder /go/bin/simple-image-server /go/bin/simple-image-server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENV ABS_ACCOUNT_KEY=YOUR_ACCOUNT_KEY
ENV ABS_ACCOUNT_NAME=YOUR_ACCOUNT_NAME
ENV ABS_AZURE_URI=YOUR_AZURE_URI
EXPOSE 8080
ENTRYPOINT ["/go/bin/simple-image-server"]
