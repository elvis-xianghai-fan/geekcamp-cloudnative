FROM golang:1.17 AS builder
ENV CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /build
COPY . .
RUN go build -o httpserver .


FROM alpine
COPY --from=builder /build/httpserver /
ENV SERVICE_PORT=8080
EXPOSE ${SERVICE_PORT}
ENTRYPOINT /httpserver
