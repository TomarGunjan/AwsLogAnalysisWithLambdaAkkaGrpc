import json
import boto3
from datetime import datetime
from datetime import timedelta
import re
import logging
import hashlib
import os

# import requests

lines = []
timeFormat = "%H:%M:%S.%f"

def getTime(line):
    content = re.split("\s+", line)
    return datetime.strptime(content[0], timeFormat)



def getIdx(lines, time, start, end):
    if(start==end):
        return start
    if(end<start):
        return start
    mid=int((start+end)/2)
    mid_time = getTime(lines[mid])
    if (time>mid_time) :
        return getIdx(lines, time, mid+1, end)
    elif (time<mid_time):
        return getIdx(lines, time, start, mid-1)
    else :
        return mid


def findMatchingMessageList(startidx, endidx, lines, pattern):
    messages=[]
    for i in range(startidx, endidx+1):
        content = re.split("\s+", lines[i])
        patternRegex = re.compile(pattern)
        print("index",i)
        print(lines[i])
        print(content)
        grps = patternRegex.search(content[5])
        
        if(grps!=None):
            result = hashlib.md5(grps.group(0).encode())
            messages.append({"time":content[0],"messageHash":str(result.digest())})
    return messages



def getLogList1(date, start_time, end_time, pattern):
    fileName = "LogFileGenerator." + date + ".log"
    s3 = boto3.resource('s3')
    bkt = s3.Bucket("mylogfiles")
    for objs in bkt.objects.all():
        if objs.key == fileName:
            content = objs.get()['Body'].read().decode('utf-8')
            lines = content.split("\n")
            startidx = getIdx(lines,start_time,0,len(lines)-2)
            endidx = getIdx(lines, end_time, 0, len(lines)-2)
            return findMatchingMessageList(startidx, endidx, lines, pattern)
    return []
    
def getLogList(date, start_time, end_time, pattern):
    fileName = "LogFileGenerator." + date + ".log"
    path = "/mnt/efs"
    for file in os.listdir(path):
        if (file==fileName):
            if os.path.isfile(os.path.join(path, file)):
                f = open(os.path.join(path, file),"r")
                content = f.read()
                lines = content.split("\n")
                startidx = getIdx(lines,start_time,0,len(lines)-2)
                endidx = getIdx(lines, end_time, 0, len(lines)-2)
                return findMatchingMessageList(startidx, endidx, lines, pattern)
    return []

def lambda_handler(event, context):
    logging.info("lambda function started")
    date = event['queryStringParameters']['date']
    time = event['queryStringParameters']['time']
    delta = float(event['queryStringParameters']['delta'])
    pattern = event['queryStringParameters']["pattern"]
    logging.info("##event received")
    logging.info(event)
    inputTime = datetime.strptime(time, timeFormat)
    startTime = inputTime-timedelta(minutes=delta)
    endTime = inputTime + timedelta(minutes=delta)
    list=getLogList(date,startTime,endTime, pattern)
    logging.info("successfully retrieved log data")
    return {
        "statusCode": 200,
        "body": json.dumps({
            "date": date,
            "logs":list
        }),
    }
