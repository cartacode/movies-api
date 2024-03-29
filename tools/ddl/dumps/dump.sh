#!/bin/bash

rm -rf vuliapi_mongo.dump

mongo vuliapi --eval "db.dropDatabase()"

mongodump --host api-mongodb.stage.vuli.io -d vuliapi --port 27017 --out vuliapi_mongo.dump

mongorestore --db vuliapi vuliapi_mongo.dump/vuliapi

rm -rf vuliapi_mongo.dump
