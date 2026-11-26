#!/bin/bash
# Valid curl
curl -X POST -H 'Content-Type: application/json' -d '{"body":"This is a test"}' http://localhost:8080/api/validate_chirp
# Empty body
curl -X POST -H 'Content-Type: application/json' -d '{"body":""}' http://localhost:8080/api/validate_chirp
# TOo long body
curl -X POST -H 'Content-Type: application/json' -d '{"body":"This is a looong message because I want to make sure that it is not valid and it will fail without a doubt...orem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua orem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"}' http://localhost:8080/api/validate_chirp
# Curse words added
curl -X POST -H 'Content-Type: application/json' -d '{"body":"Kerfuffle this is a simple test sharbert!"}' http://localhost:8080/api/validate_chirp
# admin access reset
curl -X POST  http://localhost:8080/admin/reset
# testing chirps
curl -X POST -H 'Content-Type: application/json' -d '{"body":"If you're committed enough, you can make any story work.","user_id":"ea954000-5ce2-4370-9e90-5a36a1be989f"}' http://localhost:8080/api/chirps
