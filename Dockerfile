FROM golang:alpine as builder

RUN apk add --update git

RUN mkdir -p $GOPATH/src/build

ADD . $GOPATH/src/build/

WORKDIR $GOPATH/src/build

RUN go get ./...

RUN go build -o main .

FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

COPY --from=builder /go/src/build/main /app/

COPY --from=builder /go/src/build/wisdoms.txt /app/

WORKDIR /app

CMD ["./main"]