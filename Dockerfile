ARG APP_VERSION="debug"
FROM golang:1.18-alpine3.16 as builder
COPY . ./src
WORKDIR src
RUN go build  -ldflags "-X 'main.version=${APP_VERSION}'" -o entrypoint
FROM scratch
COPY --from=builder src/entrypoit /entrypoint
EXPOSE 8080
ENTRYPOINT ["/entrypoint"]
