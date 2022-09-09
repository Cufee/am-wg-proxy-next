FROM golang:alpine as builder

# Build
WORKDIR /app 
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o binary .

# Tailscale
FROM alpine:latest as tailscale
WORKDIR /app

ENV TSFILE=tailscale_1.30.1_amd64.tgz
RUN wget https://pkgs.tailscale.com/stable/${TSFILE} && tar xzf ${TSFILE} --strip-components=1
RUN mkdir -p /var/run/tailscale /var/cache/tailscale /var/lib/tailscale
RUN apk add ip6tables iptables

# Binary
ENV TZ=Europe/Berlin
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder /app/binary .
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ./tailscaled & ./tailscale up --authkey=${TAILSCALE_AUTHKEY} --advertise-tags=tag:service --hostname=${TAILSCALE_APP_NAME}; ./binary && tailscale logout