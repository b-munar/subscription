FROM golang:1.20-bullseye AS builder

EXPOSE 80

WORKDIR /usr/src

COPY . .

RUN go build -o subscription .

FROM debian:bullseye-slim

COPY --from=builder /usr/src/subscription /usr/local/bin/subscription

CMD ["subscription"]