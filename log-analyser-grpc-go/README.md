# Log-Analyser-GRPC-GO
## GRPC framework for setting up Server and Client for fetching Logs Data

## Overview

This framework creates a server which accepts RPC requests from client and in turns interacts with Akka HTTP Server created in [previous step](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/tree/master/akka-http-loganalyser-scala#readme). Go has been used for developing this Framework. This framework also provides a client to test the server

## Pre-Requisites
1. Go should be installed and GOPATH should be setup
2. [Previous Step](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/tree/master/akka-http-loganalyser-scala#readme) should be completed and AkkaHttp Application Server should be up and running

## Project Structure

  -  server/main.go - This file contains code for setting up the server
  -  proto This folder contains proto file and pb.go(stub files)
  -  services/logService.go - This file contains implementation for service created in proto file
  -  service/logService_test.go - This file contains unit test for the service
  -  helper/helper.go - This file contains helper function to make http calls
  -  constants/constants.go - This file contains constant values
  -  client/main.go - This file contains client code to be used for testing grpc server
  -  inpuFiles/ - This directory contains input to be fed to client
  
## SetUp

1. Clone the repository into src folder of your GOPATH
2. Setting up server
    - Navigate to server folder in the project from terminal 
    - Run following command ```go run .\main.go```
    - You should see a message saying ``` server listening at [::]:50051 ```
 
## Testing

1. Open another terminal and navigate to client direcyory in project.
2. Run command to run positive scenario go run main.go pass
3. Run command to run negative scenario go run main.go fail

***Note*** The input folder contains request input which can be modified and corresponding path can be provided in client/main.go to customize the testing

### Input
  
   {
  "date": "2022-10-06",
  "time": "00:50:00.000",
  "delta": "30",
  "pattern": "(%5Ba-c%5D%5Be-g%5D%5B0-3%5D%7C%5BA-Z%5D%5B5-9%5D%5Bf-w%5D)%7B5%2C15%7D"
  }

### Output
1. Success Response Status - 200
  ```{"date": "2022-10-06","logs": [{"messageHash": "b'\\xb3+\\x9c\\x7f\\xfb3Q\\x00\\x98\\xdd\\x13c\\x16n\\xb4%'","time": "00:29:21.017" },]}```
2. Error Response Status 404
  ```{"error":"The data you are looking for is not currently present"}```
  
  
## Technology Used
1. GRPC framework
2. Golang
3. Mockery

## Navigation
  - [Go to main Project](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/README.md)
  - [Go to previous step(Step 3)](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/tree/master/akka-http-loganalyser-scala#readme)


