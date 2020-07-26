# search-nearest-places
A Go HTTP server, which exposes one endpoint /api/v1/places to get the nearest places, given location as queryParameter

If server is running on port 9080, 
then below API call will generate response like ./sample-response.json

GET http://localhost:9080/api/v1/places?location=Berlin

# use docker image

steps to follow
1.docker pull allahthedev/search-nearby-places:1.1

2.docker run -it -d -p 9080:9080 -i allahthedev/search-nearby-places:1.1

3.hit url : GET http://localhost:9080/api/v1/places?location=London
