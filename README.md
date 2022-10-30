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
  "delta": "30",      //in minutes
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



