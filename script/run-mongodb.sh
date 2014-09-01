#!/bin/sh

# See https://github.com/dockerfile/mongodb
docker run --name rima-mongodb -p 27017:27017 -d dockerfile/mongodb
