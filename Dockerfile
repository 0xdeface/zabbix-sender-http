ARG APP_VERSION="debug"
FROM golang:1.18-alpine3.16 as builder
COPY . ./src
WORKDIR src
RUN go build  -ldflags "-X 'main.version=${APP_VERSION}'" -o zabbix-http
RUN ls
RUN pwd
FROM scratch
COPY --from=builder zabbix-http /main
EXPOSE 8080
ENTRYPOINT ["/main"]
