# search-nearest-places
A Go HTTP server, which exposes one endpoint /api/v1/places to get the nearest places, given location as queryParameter

If server is running on port 9080, 
then below API call will generate response like ./sample-response.json

GET http://localhost:9080/api/v1/places?location=Berlin
