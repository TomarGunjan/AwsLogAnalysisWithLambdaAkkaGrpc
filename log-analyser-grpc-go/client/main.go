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

/** This function sets up a client to test GRPC server
The command line argument is used to determing which input file to be used. The values could
1. pass - happy case will be run
2. Fail - failure case will be run
3. Any other value - happy case will be run

*/

func main() {
	flag.Parse()
	log.Printf("setting up client")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := proto.NewLogAnalyserServiceClient(conn)
	wd, _ := os.Getwd()
	path := filepath.Join(wd, "../inputFiles/pass.txt")
	if (len(os.Args) > 1) && (os.Args[1] == "fail") {
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
