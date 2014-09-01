#!/bin/sh

docker run -it --link rima-mongodb:mongodb --rm dockerfile/mongodb mongo --host mongodb
