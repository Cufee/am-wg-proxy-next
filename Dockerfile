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

# Install Doppler CLI
RUN wget -q -t3 'https://packages.doppler.com/public/cli/rsa.8004D9FF50437357.key' -O /etc/apk/keys/cli@doppler-8004D9FF50437357.rsa.pub && \
  echo 'https://packages.doppler.com/public/cli/alpine/any-version/main' | tee -a /etc/apk/repositories && \
  apk add doppler

# Inject secrets
ENTRYPOINT ["doppler", "run", "--"]

COPY --from=builder /app/entrypoint.sh ./entrypoint.sh
CMD ./entrypoint.sh