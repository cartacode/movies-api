# Go API Server for Vuli

Vuli Movie Delivery API

## Overview
- API version: 1
- Build date: N/A
For more information, please visit [Drone Build Page](https://drone.vuli.io/VuliTv/go-movie-api)


### Running the server
To run the server, follow these simple steps:

```
# install dep & gin
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
go get github.com/codegangsta/gin

# clone this repo to $GOPATH/src/github.com/VuliTv/go-movie-api
dep ensure

# easiest to get running
docker-compose up

# for dev you need to have mongo and redis running locally or set these ENV variables
#
# example for stage:
#
# export MONGO_HOST="api-mongodb.stage.vuli.io"
# export REDIS_ADDRESS="api-redis.stage.vuli.io:6379"

gin --immediate

http://127.0.0.1:3001
```


### Documentation

##### API w/ environments and examples
[Postman Docs and Collection](https://vulitv.postman.co/collections/2363352-b39fcf84-0bf4-4245-b2a1-f55bdefc3852?workspace=b34041b8-635b-4ba7-8435-23b9999b0e86)

##### Go Model to MongoDB

[Model PDF](docs/models.pdf)
