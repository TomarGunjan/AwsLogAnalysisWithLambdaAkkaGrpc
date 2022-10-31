# Deploying Modified Log Generator on EC2 and its integration with EFS

## Overview
This repository is Step 1 of Homework 2 Project AwsLogAnalysisWithLambdaAkkaGrpc and describes changes in Log Generator and its deployment

## About EC2

Amazon Elastic Compute Cloud (Amazon EC2) provides scalable computing capacity in the Amazon Web Services (AWS) Cloud. Using Amazon EC2 eliminates your need to invest in hardware up front, so you can develop and deploy applications faster. You can use Amazon EC2 to launch as many or as few virtual servers as you need, configure security and networking, and manage storage. 

## About EFS

Amazon Elastic File System (Amazon EFS) provides a simple, serverless, set-and-forget elastic file system for use with AWS Cloud services and on-premises resources. It is built to scale on demand to petabytes without disrupting applications, growing and shrinking automatically as you add and remove files, eliminating the need to provision and manage capacity to accommodate growth

We are going to mount EFS on our EC2 instance and utilize it.

## Pre-requisites
- Valid AWS Account

## Setup

1. Creating a EFS
    - Create Default VPC and default Security Group
    - Navigate to EFS
    - Click on Create a File System
    - Select default vpc and click on Create
    - Once created navigate to created EFS instance and to access points section
    - Create Default Access Point
    - Select the Security Group created in 1st step under Networks tab and save all configs
 
2. Creating EC2 instance
     - Create a Security group with HTTP, HTTPS, RDP AND SSH rules
    - Create key pair and save it
    - Navigating to EC2 section
    - Click on Launch Instance
    - Select Amazon Linux as AMI
    - Select the created key pair
    - Under Network Settings select both security groups created in 1st and 2nd step
    - Click on Launch Instance
  
 3. Setting up Log Generator in EC2 instance
    - Select the created instance in previous step and click on connect
    - Use the info provided under Connect section to connect to create instance from your system
    - Use following commands to install java 11 and sbt
      - sudo amazon-linux-extras install java-openjdk11
      - sudo rm -f /etc/yum.repos.d/bintray-rpm.repo
      - curl -L https://www.scala-sbt.org/sbt-rpm.repo > sbt-rpm.repo
      - sudo mv sbt-rpm.repo /etc/yum.repos.d/
      - sudo yum install sbt
    - using winscp or git copy the [log generator](https://github.com/TomarGunjan/AwsLogAnalysisWithLambdaAkkaGrpc/tree/master/ModifiedLogGenerator/LogFileGenerator) to your EC2 instance
    - Set up a cronscript in your EC2 instance to run project periodically
        - create a script cronscript.sh and provide command "cd (to LogGenerator Project); sbt clean compile run;
        - keep the script in home folder (ec2-user)
        - Type crontab -e and set cron expression ex 00 00 * * * path to cronscript 
 
 4. Mounting EFS to EC2
    - Run following command in your instance to get Amazon-EFS Tools
      - sudo yum install amazon-efs-utils
    - Navigate to Log Generator directory and create a folder with name "efs"
    - Navigate to EFS created in Step and click on Attach
    - Copy command given under EFS mount helper and run it on your EC2 instance inside Log Generator project folder
    - Change permissions for efs folder to 777

 ## Changes in Log Generator Project
 
 1. Changes to logback-text.xml has been made to create logs under efs folder

## Sources

- [AWS Guide](https://aws.amazon.com/)
- [LogGenerator Project by Prof. Mark Grechanik](https://github.com/TomarGunjan/CS441_Fall2022/tree/main/LogFileGenerator)
