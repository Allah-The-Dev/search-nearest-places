# search-nearest-places
A Go HTTP server, which exposes one endpoint /api/v1/places to get the nearest places, given location as queryParameter

If server is running on port 9080, 
then below API call will generate response like ./sample-response.json

To run this as a Go program steps are:

1. Clone the repo

2. Inside repo folder run command :: go run .

GET http://localhost:9080/api/v1/places?location=Berlin

# use docker image

steps to follow

1. docker pull allahthedev/search-nearby-places:1.1

2. docker run -it -d -p 9080:9080 -i allahthedev/search-nearby-places:1.1

3. hit url : GET http://localhost:9080/api/v1/places?location=London

# Key information of project

1. Here API calls are made parallel using goroutine and sync package

2. Unit test is written using table driven test cases in Golang

3. It have 88.4% of unit test coverage

4. It uses LRU cache

5. Docker image size is highly optimized with size of 13.8 MB, 

6. Docker file is a multistage file


