FROM golang:latest AS builder
WORKDIR /app
COPY .. .
RUN chmod +x /app/bootstrap.sh
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/app ./
COPY --from=builder /app/bootstrap.sh ./

CMD ["/bin/sh", "/app/bootstrap.sh"]
