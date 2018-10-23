FROM golang:alpine

RUN apk add git curl
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN mkdir -p /go/src/github.com/VuliTv/api/
WORKDIR /go/src/github.com/VuliTv/api/
COPY . .
RUN dep ensure


EXPOSE 3001

CMD [ "go", "run", "main.go"]
