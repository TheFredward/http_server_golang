#!/bin/bash
# Valid curl
curl -X POST -H 'Content-Type: application/json' -d '{"body":"This is a test"}' http://localhost:8080/api/validate_chirp
# Empty body
curl -X POST -H 'Content-Type: application/json' -d '{"body":""}' http://localhost:8080/api/validate_chirp
# TOo long body
curl -X POST -H 'Content-Type: application/json' -d '{"body":"This is a looong message because I want to make sure that it is not valid and it will fail without a doubt...orem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua orem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"}' http://localhost:8080/api/validate_chirp
