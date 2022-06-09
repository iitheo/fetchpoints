## What is this repository for? ##
* Fetch Rewards Points Service

## What is in this README?
* Url for Live App
* How do I set up app locally? 
* Postman documentation


## How do I set up app locally? ##
There are 2 ways:
Method 1: Open your terminal and navigate to the folder where main.go is located
cmd->server
1. Next type the command go run main.go
2. This will start the app in your local pc at port 8081
3. Next import the postman collection and call the endpoints.

Method 2: Docker
1. Ensure you have docker installed on your computer.
2. Start up docker on your computer and ensure it is running.

## Steps to Start App ##
### Next build your docker image. ###

To build your docker image:
1. Open your Terminal. Clone the app from the github repo by running the command below.
2. git clone https://github.com/iitheo/fetchpoints.git
3. Navigate to the root path of the app. This is the same path as the Dockerfile. Run ls in your terminal and verify you can see the Dockerfile
4. Build the docker image by running the command below:
5. docker build -t myfetchrewardsapp:1 .
6. This will build the docker image.
7. Next run the docker container by running the command below:
8. docker run -dp 8081:8081 myfetchrewardsapp:1
9. Now, you have the app running.
10. Open your browser or Postman.
11. type this localhost:8081/v1/points/getall
12. You will get a response like this or something similar:


{
"success": true,
"data": [
{
"payer": "MILLER COORS",
"points": 10000,
"timestamp": "2020-11-01T14:00:00+01:00"
},
{
"payer": "DANNON",
"points": 300,
"timestamp": "2020-10-31T10:00:00+01:00"
},
{
"payer": "DANNON",
"points": 1000,
"timestamp": "2020-11-02T14:00:00+01:00"
},
{
"payer": "UNILEVER",
"points": 200,
"timestamp": "2020-10-31T11:00:00+01:00"
},
{
"payer": "DANNON",
"points": 250,
"timestamp": "2020-10-31T15:00:00+01:00"
}
],
"message": "5 record(s) successfully fetched"
}


CONGRATULATIONS!!!

### The following documents are provided: ###

### Postman documentation ###
Postman link
https://www.getpostman.com/collections/7bb87828d2cac7b3849b

Postman Documentation
3 endpoints are listed together with sample request and response.

They are:
1. Get All Points
2. Post Add Points
3. Post Spend Points
