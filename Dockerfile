FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN go build .

EXPOSE 8866

ENTRYPOINT ["./tcp-pack"]