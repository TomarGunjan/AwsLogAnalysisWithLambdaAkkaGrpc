# AwsLogAnalysisWithLambdaAkkaGrpc
### Repository to solve Homework 2 with Lambda Function AkkaHttp Rest Framework with Scala and Grpc Framework with Go
Submitted By - Gunjan Tomar

UIN - 674659382

Email Id : gtomar2@uic.edu

## Overview
This objective of this repository is to analyse log files and fetch logs that matches a certain pattern with a certain time interval

### Input parameters

{
  "date": "2022-10-06",
  "time": "00:00:00.000",
  "delta": "30",
  "pattern" : "abcd"
}

### Output

{ "date": "2022-10-06",
    "logs": [
        {
            "messageHash": "b'\\xb3+\\x9c\\x7f\\xfb3Q\\x00\\x98\\xdd\\x13c\\x16n\\xb4%'",
            "time": "00:29:21.017"
        }]
}

## Steps for the project

1. Modified Log Generator Project (Original by by Prof. Mark Grechanik) Deployment on EC2 and its integration with AWS EFS
2. Setting up Lambda Service with API Gateway an its integration with AWS EFS
3. Setting up Akka Http Server connecting to Lambda Services
4. Setting up GRPC server and client with server connecting to Akka Http Server




