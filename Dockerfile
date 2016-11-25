FROM golang:alpine
MAINTAINER sklirg

ENV GIN_MODE=release
ENV PORT=8000
ENV APP=/go/src/github.com/sklirg/wcgstats-api/
EXPOSE 8000

RUN apk add --no-cache --update git

RUN mkdir -p $APP
WORKDIR $APP

COPY . .
RUN go get
RUN go build

CMD ["./wcgstats-api"]
