FROM spryhq/golang-dep:alpine-1.10

RUN mkdir -p /usr/share/dict
RUN wget https://s3.amazonaws.com/vuli-public-assets/words -O /usr/share/dict/words
RUN mkdir -p /go/src/github.com/VuliTv/go-movie-api/
WORKDIR /go/src/github.com/VuliTv/go-movie-api/
COPY . .
RUN dep ensure


EXPOSE 3001

CMD [ "go", "run", "main.go"]
