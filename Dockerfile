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

# Runner
FROM alpine:latest as runner

# Tailscale
RUN apk add ip6tables iptables
COPY --from=tailscale /app/tailscale .
COPY --from=tailscale /app/tailscaled .
RUN mkdir -p /var/run/tailscale /var/cache/tailscale /var/lib/tailscale

# Binary
ENV TZ=Europe/Berlin
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder /app/binary .
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Run entrypoint script
COPY --from=builder /app/entrypoint.sh .
RUN chmod +x entrypoint.sh
ENTRYPOINT [ "./entrypoint.sh" ]