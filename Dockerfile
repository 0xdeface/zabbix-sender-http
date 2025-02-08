ARG APP_VERSION="debug"
ARG ENTRYPOINT="entrypoint"
FROM golang:1.18-alpine3.16 as builder
COPY . ./src
WORKDIR src
RUN go build  -ldflags "-X 'main.version=${APP_VERSION}'" -o ${ENTRYPOINT}
FROM scratch
COPY --from=builder src/${ENTRYPOINT} /${ENTRYPOINT}
EXPOSE 8080
ENTRYPOINT ["/${ENTRYPOINT}"]
