# Lambda Functions for Chceking if Log Exists and fetching the logs

## AWS Lambda

AWS Lambda is a serverless, event-driven compute service that lets you run code for virtually any type of application or backend service without provisioning or managing servers. You can trigger Lambda from over 200 AWS services and software as a service (SaaS) applications, and only pay for what you use.

### AWS Lambda Trigger with API Gateway Trigger
  
API Gateway provides tools for creating and documenting web APIs that route HTTP requests to Lambda functions.

### AWS Lambda with EFS as file system

We are using EFS as file system created in Step 1 as file system for AWS Lambda

## Technology used

- AWS Lambda
- AWS Elastic File Storage(EFS)
- AWS API Gateway Trigger

## What are we doing?
 
We are creating lambda functions and selecting EFS as file system for them so that both Log Generator in EC2 instance and Lambda function would be able to access a common space in efs. Log Generator would be creating data and AWS Lambda would be running function son that data. We are also setting AWS API Gateway as trigger for Lambda functions so that they can be accesses using HTTP endpoints

## SetUp

1. Creating IAM role for AWS Lambda
    - Navigate to IAM -> Access Management -> Roles
    - Click on Create Role
    - Select Lambda under "Use Case" and click on Next
    - Select following Policies for your rule and click on Next
      - AWSLambdaExecute
      - AWSLambdaVPCAccessExecutionRole
      - AmazonElasticFileSystemClientFullAccess
    - Click on Create Role
2. Creating Lambda Function
    - Navigate to AWS Lambda
    - Click on Create Function.
    - Provide Function Name , Select Runtime as python 3.8
    - Click on "Change Default Execution Role
    - Select "Use an existing role" and select role created in step 1
    - Setting up VPC and File System
      - Click on Advanced Settings
      - Click on Enable VPC
      - Select VPC created while setting up [EFS in Log generator](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/ModifiedLogGenerator/README.md)
      - Choose 2 subnets including the area where you are creating your lambda function
      - Select Default Security group created during setup of [EFS in Log generator](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/ModifiedLogGenerator/README.md)
    - Click on Create Function
3. Setting up File System for Lambda Function
    - Navigate to details of created function
    - Go to Configuration -> File Systems
    - Click on Add File System
    - Select File System Created while setting up [Log Generator](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/ModifiedLogGenerator/README.md)
    - Select available Access point(created while setting up Log Generator)
    - provide local mount path values as /mnt/efs
    - Click on Save
4. Uploading Lambda Function Code
    - Clone the code provided under(GetLogs/CheckLogs) [lambdas directory](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/tree/master/lambdas)
    - zip the file app.py 
    - Go to details of your Lambda Function and navigate to code and click on Upload From on Top Right corner
    - Upload the zip file created and your directory should reflect app.py file
5. Creating api Gateway
    - Navigate to Amazon API Gateway and click on Create Api
    - Click on Build option under Rest API 
    - Provide API Name and click on Create API
    - Setting up Get End point under created api
    - Navigate to your create api and click on it. Click on Actions-> Create Resource
    - Provide appropriate resource name and resource path and click on Create Resource
    - Select the created Resource and Click on Actions -> Create Method
    - Select GET as method
    - Select "Integration Type" as "Lambda Function" , your Lambda region, and Type in the name of your Lambda function. it should come as option.
    - Select your Lambda Function and click on Save
    - Select your api and click on Actions-> Deplpy API
    - Select New stage and provide name as "stage"
 6. Creating API Trigger for Lambda Function
    - Navigate to Lambda Function Created in Step 2 -> Triggers
    - Click on Add Trigger
    - Select "API Gateway" as trigger
    - Select "Use existing API" and select API created in step 5
    - Select "Deployment Stage" as stage
    - Click on Add
 7. Follow steps 2 to 6 for creating second lambda Function
CloudWatch
## Testing

1. AWS Testing 
    - Click on Test Tab under Lambda Function and provide following payload under Event Json
    {
  "queryStringParameters": {
    "date": "2022-10-30",
    "time": "20:04:00.000",
    "delta": "30",
    "pattern": "([a-c][e-g][0-3]|[A-Z][5-9][f-w]){5,15}"
  }
}
    - Click on Test button. You can also check Logs under Monitor Tab which takes us to CloudWatch
    
 2. Testing using API Gateway 
    - Navigate to Configurations->Triggers
    - You will find the endpoint for the api there
    - The lambda function can be tested using Postman or curl. A sample curl for checklogs api is following
    curl --location --request GET 'https://peh6k1md85.execute-api.us-east-1.amazonaws.com/Prod/logcheck?time=13:00:00.000&date=2022-09-27&delta=30'
    
 ## API Input and Output
 
 # Input
 
 1. Check Logs Api
   - takes follwoing query parameter
         1. time
         2. date
         3. delta

2. Get Logs Api
  - takes following query parameter
        1. time
        2. delta
        3. date
        4. pattern
  - **Note** - pattern should be encoded when hitting the endpoint from postman. Just right click on pattern string and select encode

# Output

1. Check Logs Api '
    - provides a flag to tell if there are any logs entry present for input time interval
        ``` {"isPresent": true} ```
2. Get Logs API
    - provides date for the query log, and a list containing 1. md5 has for matched string and its timestamp
      ```{
    "date": "2022-10-06",
    "logs": [
        {
            "messageHash": "b'\\xb3+\\x9c\\x7f\\xfb3Q\\x00\\x98\\xdd\\x13c\\x16n\\xb4%'",
            "time": "00:29:21.017"
        },]}```
        
## What's next?
- [Go to main Project](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/README.md)
- [Go to Next Step(Step 3)](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/blob/master/lambdas/README.md)
- [Go to Previous Step(Step 1)](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/edit/master/ModifiedLogGenerator/README.md)
