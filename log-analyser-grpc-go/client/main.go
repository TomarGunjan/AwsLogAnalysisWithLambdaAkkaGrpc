package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
	"log"
	"log-analyser-grpc-go/proto"
	"os"
	"path/filepath"
	"time"
)

var (
	serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewLogAnalyserServiceClient(conn)
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "../inputFiles/pass.txt")
	if len(os.Args) > 1 {
		path = filepath.Join(wd, "../inputFiles", os.Args[1])
	}
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var request *proto.GetMessageInIntervalRequest
	json.Unmarshal([]byte(byteValue), &request)
	response, err := client.GetMessageInInterval(ctx, request)
	if err != nil {
		print(err.Error())
	}
	log.Print(response)
}
