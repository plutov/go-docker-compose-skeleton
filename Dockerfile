FROM golang:1.8

ADD . /go/src/github.com/plutov/go-docker-compose-skeleton
WORKDIR /go/src/github.com/plutov/go-docker-compose-skeleton

RUN curl --silent https://glide.sh/get | sh
RUN go get bitbucket.org/liamstask/goose/cmd/goose
RUN go get github.com/go-swagger/go-swagger/cmd/swagger

ENTRYPOINT bash -c "/go/src/github.com/plutov/go-docker-compose-skeleton/run.sh"