FROM golang:1.22 AS builder

WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /app ./cmd/maldon/

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app /app


EXPOSE 8080

ENTRYPOINT ["/app"]
