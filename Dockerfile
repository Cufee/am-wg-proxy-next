FROM golang:alpine as builder

# Build
WORKDIR /app 
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o binary .

# Runner
FROM scratch as runner
WORKDIR /app

ENV TZ=Europe/Berlin
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder /app/binary /usr/bin/
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "binary" ]