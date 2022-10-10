FROM golang:alpine as builder

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o binary .

ADD https://github.com/ahmetb/runsd/releases/download/v0.0.0-rc.15/runsd /app/runsd
RUN chmod +x /app/runsd

FROM scratch

ENV TZ=Europe/Berlin
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder /app/runsd /bin/runsd
COPY --from=builder /app/binary /
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["runsd", "--", "/binary"]