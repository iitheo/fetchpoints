## What is this repository for? ##
* Fetch Rewards Points Service

## What is in this README?
* Url for Live App
* How do I set up app locally? 
* Postman documentation
* How to run unit test
* What I would have done differently if it were Production API


## How do I set up app locally? ##
## Requirements: ##
1. Ensure you have docker installed on your computer.

2. Start up docker on your computer and ensure it is running.

## Steps to Start App ##
### Next build your docker image. ###

To build your docker image:
1. Open your Terminal. Clone the app from the github repo by running the command below.
2. git clone https://github.com/iitheo/mypokemon.git
3. Navigate to the root path of the app. This is the same path as the Dockerfile. Run ls in your terminal and verify you can see the Dockerfile
4. Build the docker image by running the command below:
5. docker build -t myfetchrewardsapp:1 .
6. This will build the docker image.
7. Next run the docker container by running the command below:
8. docker run -dp 8081:8081 myfetchrewardsapp:1
9. Now, you have the app running.
10. Open your browser or Postman.
11. type this localhost:8081/v1/pokemon/get/misdreavus
12. You will get a response like this or something similar:

{
"data": {
"name": "misdreavus",
"description": "It likes playing\nmischievous tricks\nsuch as screaming\fand wailing to\nstartle people at\nnight.",
"habitat": "cave",
"isLegendary": false,
"translatedDescription": "Playing mischievous tricks such as screaming and wailing to startle people at night,  it likes."
},
"message": "data successfully fetched",
"status": true
}


CONGRATULATIONS!!!

### The following documents are provided: ###

### Postman documentation ###
https://www.getpostman.com/collections/02ec0f5ec7b0e4833f7b

Postman Documentation
6 endpoints are listed together with sample request and response.

The first 5 endpoints provide context on the pokemon api and also
provides sample data to test the MyPokemon App.

The custom endpoint for getting fun description from MyPokemon App is:
GET Fun Pokemon Data

##  How to run unit test in the app
Just one unit test was written because the other tests would require communicating across the network. And this falls outside a unit test.
Navigate to the directory 'tests' it is on the same folder hierarchical level with cmd and pkg
Then run go test -run . -v


##  What I would have done differently if it were Production API
* I would ask more questions on the functional requirements and the different users of the app and use cases of the app. This information
would guide my decision on non-functional requirements like scaling, availability, latency, security, etc.
* I would use redis to cache more data to save API calls to both Pokemon and the Translation APIs.
* I would store secrets and other keys like the redis connection string in an environment variable instead of committing to code.
