FROM golang:1.18-alpine3.16 as builder
ARG APP_VERSION="debug"
COPY . /code
WORKDIR /code
RUN go build  -ldflags "-X 'main.version=${APP_VERSION}'" -o entry
FROM scratch
COPY --from=builder /code/entry /entry
EXPOSE 8080
ENTRYPOINT ["/entry"]
