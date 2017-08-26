# build stage
FROM golang:1.8 AS build-env

#RUN apk update
#RUN apk add git
#RUN apk add g++
#RUN apk add flite-dev
RUN apt-get update && apt-get install -y flite-dev
#ADD . /src
RUN mkdir /go/src/app
COPY main.go /go/src/app
COPY vendor /go/src/vendor

#RUN cd /src && go get -d -v ./... && go build -o goapp
#RUN cd /src && go build -o goapp
RUN go install app

# final stage
FROM debian:jessie-slim

#RUN apk update && apk add flite-dev
RUN apt-get update && apt-get install -y flite-dev
#WORKDIR /bin/app
COPY --from=build-env /go/bin/app /app
ENTRYPOINT ["/app"]