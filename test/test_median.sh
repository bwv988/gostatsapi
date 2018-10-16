#!/bin/sh

# Should be 2.
curl -X POST -d '{"data": ["0", "2", "4", "6", "4", "2", "2", "2", "1"]}' http://localhost:8080/api/v1/median
