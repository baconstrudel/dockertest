FROM golang:alpine AS build-env
ADD . /go/src/github.com/baconstrudel/dockertest
WORKDIR /go/src/github.com/baconstrudel/dockertest
RUN go build -o dockertest

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/github.com/baconstrudel/dockertest/dockertest /app/
EXPOSE 8080
ENTRYPOINT ./dockertest