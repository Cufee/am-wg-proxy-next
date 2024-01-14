FROM golang:1.21 as build

WORKDIR /workspace
COPY . ./

# add go modules lockfiles
RUN go mod download

# install task
RUN go install github.com/go-task/task/v3/cmd/task@latest

# generate the Prisma Client Go client
RUN task build:docker

FROM scratch as run

ENV TZ=Europe/Berlin
ENV ZONEINFO=/zoneinfo.zip
COPY --from=build /usr/local/go/lib/time/zoneinfo.zip /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build /workspace/app .

CMD ["./app"]