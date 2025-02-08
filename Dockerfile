FROM golang:1.18-alpine3.16 as builder
ARG APP_VERSION="debug"
ARG ENTRYPOINT="entrypoint"
COPY . ./src
WORKDIR src
RUN echo $ENTRYPOINT
RUN echo ${APP_VERSION}
RUN go build  -ldflags "-X 'main.version=${APP_VERSION}'" -o ${ENTRYPOINT}
FROM scratch
ARG APP_VERSION="debug"
ARG ENTRYPOINT="entrypoint"
COPY --from=builder src/${ENTRYPOINT} /${ENTRYPOINT}
EXPOSE 8080
ENTRYPOINT ["/${ENTRYPOINT}"]
