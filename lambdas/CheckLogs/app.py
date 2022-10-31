import json
import boto3
from datetime import datetime
from datetime import timedelta
import re
import logging

# import requests

lines = []
timeFormat = "%H:%M:%S.%f"

def getTime(line):
    cnts = re.split("\s+", line)
    return datetime.strptime(cnts[0], timeFormat)

def checkTimeBound(lines, startTime, endTime):
    firstLogTime = getTime(lines[0])
    lastLogTime = getTime(lines[len(lines)-2])
    return (
            (
                (startTime<=firstLogTime)and(firstLogTime<=endTime)
            )or
            (
                (startTime<=lastLogTime)
            )
    )


def message_exist(date, start_time, end_time):
    fileName = "LogFileGenerator." + date + ".log"
    path = "/mnt/efs"
    for file in os.listdir(path):
        if (file==fileName):
            if os.path.isfile(os.path.join(path, file)):
                f = open(os.path.join(path, file),"r")
                content = f.read()
                lines = content.split("\n")
                return checkTimeBound(lines, start_time, end_time)

    return False
    
def message_exist1(date, start_time, end_time):
fileName = "LogFileGenerator."+date+".log"
s3 = boto3.resource('s3')
bkt = s3.Bucket("mylogfiles")
for objs in bkt.objects.all():
    if objs.key==fileName:
        content = objs.get()['Body'].read().decode('utf-8')
        lines = content.split("\n")
        return checkTimeBound(lines, start_time, end_time)

return False



def lambda_handler(event, context):
    logger.info("lambda function LogCheck started")
    date = event["queryStringParameters"]['date']
    time = event["queryStringParameters"]['time']
    delta = float(event["queryStringParameters"]['delta'])
    logger.info("event received")
    logger.info(event)
    inputTime = datetime.strptime(time, timeFormat)
    startTime = inputTime-timedelta(minutes=delta)
    endTime = inputTime + timedelta(minutes=delta)
    found=message_exist(date,startTime,endTime)
    logger.info("successfully retrieved data")
    return {
        "statusCode": 200,
        "body": json.dumps({
            "isPresent": found,
        }),
    }

