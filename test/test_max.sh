#!/bin/sh

curl -X POST -d '{"data": ["12", "15.1", "7", "-8", "99", "0"]}' http://localhost:8080/api/v1/max/2
#curl -X POST -d '{"data": []}' http://localhost:8080/api/v1/max/2