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

1. [Modified Log Generator Project (Original by by Prof. Mark Grechanik) Deployment on EC2 and its integration with AWS EFS](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/ModifiedLogGenerator/README.md)
2. [Setting up Lambda Service with API Gateway an its integration with AWS EFS](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/lambdas/README.md)
3. [Setting up Akka Http Server connecting to Lambda Services](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/akka-http-loganalyser-scala/README.md)
4. [Setting up GRPC server and client with server connecting to Akka Http Server](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/log-analyser-grpc-go/README.md)

We will first set up a Log Generator project in EC2 instance and connect to EFS where the generated logs would be stored. We will then create Lambda Function and connect them to same access so they can run analytics on generated logs. We will then expose these lambda functions using API Gateway. We can connect to the gateway APIs using any REST Client. We will then create a Akka Http Application which would interact with gateway APIs and expose a endpoint . We will then create a grpc server which will connect exposed endpoint of AKKA HTTP application. We will Lastly create a GRPC Client to test GRPC server


### Technologies and Languagesused

- [Modified Log Generator Project (Original by by Prof. Mark Grechanik)](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/ModifiedLogGenerator/README.md)
- AWS EC2
- AWS EFS
- AWS IAM
- AWS LAMBDA with Python
- AWS API GATEWAY
- Akka Http Framework with Scala
- GRPC Framework with GOLANG

## Youtube Video

https://www.youtube.com/watch?v=_2dwsv7h30o

