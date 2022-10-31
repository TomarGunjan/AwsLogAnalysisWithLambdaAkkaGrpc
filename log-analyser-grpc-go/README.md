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
    - You should see a message saying Your project is up and running


