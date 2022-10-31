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

