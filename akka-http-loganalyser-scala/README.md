# Step 3 - Rest Akka Http Framework to make Request AWS Lambda Gateway APIs

Submitted By - Gunjan Tomar

UIN - 674659382

Email Id : gtomar2@uic.edu

## Overview

To Setup AkkaHttp Server to interact with AWS Lambda APIs

## Pre Requisites

1. Install SBT
2. Install Java 11
3. [Previous step](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/lambdas/README.md) should be complete and Lambda Services should be up on AWS

## What are we doing?
 
We are setting up a AkkaHttp server in Scala. We would be running this on ou machine. This server exposes Get endpoint . We would be using Postman or some other Rest Client to interact with it. This server acts as a client for Lambda API Gateway and interact with GetLogs and CheckLogs api created in [previous Step](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/lambdas/README.md)

## About The Project

### Important scala classes
  - logFinderApp - entry point to run the server
  - LogRoutes - Routes the request to corresponding endpoint functions
  - HttpCalls - contains utility functions and connects with Lambda api
  
### Logging
  
logback.xml file contains setup for log and use TimeBasedRolling policy. The rollover happens everyday at midnight

##Setup

1. Clone the Repository to your system
2. Make sure to add AWS Lambda APIs url in application.conf file

### Running from Local
1. cd into the cloned repository from terminal
2. Run following command ```sbt clean compile run```
3. You will see following message "Server online at http://127.0.0.1:8080/"

### Running from Intellij
1. Open the project in intellij
2. Run main function in logFinderApp class
3. You will see following message "Server online at http://127.0.0.1:8080/"

## Testing

The Http calls can be made using Postman or curl. Below is the sample curl for the request
curl --location --request GET 'http://127.0.0.1:8080/getLogs?time=00:50:00.000&date=2022-10-06&delta=30&pattern=(%5Ba-c%5D%5Be-g%5D%5B0-3%5D%7C%5BA-Z%5D%5B5-9%5D%5Bf-w%5D)%7B5%2C15%7D'

### Input 
1. Below Query Parameters should be present in the request
    - date
    - time
    - pattern
    - delta
    ***Note***- pattern should be concoded when hitting endpoint from postman. Right click query param and click on encode
    
### Output

1. Success Response Status - 200
  ```{"date": "2022-10-06","logs": [{"messageHash": "b'\\xb3+\\x9c\\x7f\\xfb3Q\\x00\\x98\\xdd\\x13c\\x16n\\xb4%'","time": "00:29:21.017" },]}```
2. Error Response Status 404
  ```{"error":"The requested resource could not be found but may be available again in the future"}```
  
## What's Next?
- [Go to main Project](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc)
- [Go to Final Step(Step 4)](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/akka-http-loganalyser-scala/README.md)
- [Go to Previous Step(Step 2)](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/tree/master/lambdas)

## Youtube video

https://www.youtube.com/watch?v=_2dwsv7h30o
