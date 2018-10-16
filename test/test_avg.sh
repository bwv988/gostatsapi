#!/bin/sh

curl -X POST -d '{"data": ["0", "2", "4", "6"]}' http://localhost:8080/api/v1/avg
