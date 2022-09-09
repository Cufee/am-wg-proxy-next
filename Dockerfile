FROM golang:alpine as builder

# Build
WORKDIR /app 
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o binary .

# Tailscale
FROM alpine:latest as tailscale
WORKDIR /app
COPY . ./
ENV TSFILE=tailscale_1.30.1_amd64.tgz
RUN wget https://pkgs.tailscale.com/stable/${TSFILE} && tar xzf ${TSFILE} --strip-components=1
COPY . ./
RUN mkdir -p /var/run/tailscale /var/cache/tailscale /var/lib/tailscale

FROM scratch as runner

WORKDIR /app

# Copy over tailscale
COPY --from=tailscale /app/tailscale .
COPY --from=tailscale /app/tailscaled .
COPY --from=tailscale /var/lib/tailscale /var/lib/tailscale 
COPY --from=tailscale /var/cache/tailscale /var/cache/tailscale
COPY --from=tailscale /var/run/tailscale /var/run/tailscale

ENV TZ=Europe/Berlin
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder /app/binary .
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ./tailscaled & ./tailscale up --authkey=${TAILSCALE_AUTHKEY} --advertise-tags=tag:service --hostname=${TAILSCALE_APP_NAME}; ./binary && tailscale logout