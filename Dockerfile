FROM golang:alpine AS builder

WORKDIR /usr/src/app

COPY . .
RUN go build

FROM alpine

WORKDIR /app

COPY --from=builder /usr/src/app/testCalcServ .
ENTRYPOINT ["/app/testCalcServ"]
