# Get (pagination and queries)
curl "http://localhost:8080/song"

curl "http://localhost:8080/song/1"
curl "http://localhost:8080/song/2"

# Post
curl \
    -X POST \
    -H "Content-Type: application/json" -d \
    '{"song1": "song", "group": "group1"}' \
	"http://localhost:8080/song"

curl \
    -X POST \
    -H "Content-Type: application/json" -d \
    '{"song2": "song", "group": "group2"}' \
	"http://localhost:8080/song"

# Delete
curl -X DELETE "http://localhost:8080/song/1"
curl -X DELETE "http://localhost:8080/song/2"

# Put
curl \
    -X PUT \
    -H "Content-Type: application/json" -d \
    '{"song1_upd": "song", "group": "group1_upd"}' \
	"http://localhost:8080/song/1"

curl \
    -X PUT \
    -H "Content-Type: application/json" -d \
    '{"song2_upd": "song", "group": "group1_upd"}' \
	"http://localhost:8080/song/2"
