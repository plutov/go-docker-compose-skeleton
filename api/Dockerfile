FROM golang:1.9

ADD . /go/src/github.com/plutov/go-docker-compose-skeleton/api
WORKDIR /go/src/github.com/plutov/go-docker-compose-skeleton/api

RUN curl --silent https://glide.sh/get | sh
RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go get github.com/go-swagger/go-swagger/cmd/swagger

ENTRYPOINT bash -c "/go/src/github.com/plutov/go-docker-compose-skeleton/api/run.sh"
