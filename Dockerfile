FROM golang:latest as builder

WORKDIR /go/src

COPY ./ ./
COPY go.mod ./
RUN go mod download

RUN go build -o /go/bin/main -ldflags '-s -w'

FROM scratch as runner
COPY --from=builder /go/bin/main /app/main

ENTRYPOINT ["/app/main"]