FROM golang:1.22.5-alpine AS builder

COPY . /codezone/source/
WORKDIR /codezone/source/

RUN go build -o ./bin/codezone cmd/api/main.go

FROM alpine:3.13

WORKDIR /root/
COPY --from=builder /codezone/source/bin/codezone .
COPY --from=builder /codezone/source/prod.env .

CMD ["./codezone", "--env-file=prod.env"]