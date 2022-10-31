package services

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"log-analyser-grpc-go/constants"
	"log-analyser-grpc-go/helper"
	"log-analyser-grpc-go/proto"
	"strings"
)

type LogAnalyserService struct {
	proto.UnimplementedLogAnalyserServiceServer
	helper helper.Helper
}

func (l *LogAnalyserService) GetMessageInInterval(ctx context.Context, in *proto.GetMessageInIntervalRequest) (*proto.GetMessageInIntervalResponse, error) {
	url := constants.BaseUrl + constants.GetLogsUrl
	url = strings.Replace(url, constants.Value1, in.Date, 1)
	url = strings.Replace(url, constants.Value2, in.Time, 1)
	url = strings.Replace(url, constants.Value3, in.Delta, 1)
	url = strings.Replace(url, constants.Value4, in.Pattern, 1)
	//resp, err := http.Get(url)
	resp, err := l.helper.MakeHttpCall(ctx, url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	println(string(body))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}
	var out *proto.GetMessageInIntervalResponse
	err = json.Unmarshal(body, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func NewLogAnalyserService(helper helper.Helper) *LogAnalyserService {
	return &LogAnalyserService{
		helper: helper,
	}
}
