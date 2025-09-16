FROM golang:1.24 AS builder

WORKDIR /heimdall

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/heimdall ./cmd/heimdall


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /heimdall/bin/heimdall ./heimdall

COPY --from=builder /heimdall/ .

EXPOSE 8080

ENTRYPOINT ["./heimdall"]

CMD ["--config","config.json"]

