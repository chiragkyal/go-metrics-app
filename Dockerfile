FROM golang:alpine3.14 as builder
RUN mkdir -p /go/src/folder
WORKDIR /go/src/folder
COPY ./ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o app .


FROM registry.access.redhat.com/ubi8/ubi-minimal:8.4
WORKDIR /
COPY --from=builder /go/src/folder/app .
CMD ["/app"]