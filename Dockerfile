FROM golang:alpine as builder

# SSH setup
ARG SSH_PRIVATE_KEY
RUN mkdir -p /root/.ssh
RUN --mount=type=secret,id=ssh_priv,dst=/id_rsa cat /id_rsa > /root/.ssh/id_rsa
RUN --mount=type=secret,id=ssh_pub,dst=/id_rsa.pub cat /id_rsa.pub > /root/.ssh/id_rsa.pub
RUN chmod 600 /root/.ssh/id_rsa

RUN apk add --update --no-cache openssh
RUN ssh-keyscan -H github.com >> /root/.ssh/known_hosts

RUN apk add git
# Override go get to use ssh instead of https
RUN git config --global url."ssh://git@github.com/".insteadOf "https://github.com/"

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o binary .

FROM scratch

ENV TZ=Europe/Berlin
ENV ZONEINFO=/zoneinfo.zip
COPY --from=builder /app/binary /usr/bin/
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["binary"]