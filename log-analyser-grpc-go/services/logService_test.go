package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"log-analyser-grpc-go/helper/mocks"
	"log-analyser-grpc-go/proto"
	"net/http"
	"testing"
)

func Test_GetMessageInInterval_Success(t *testing.T) {
	helper := &mocks.Helper{}
	service := NewLogAnalyserService(helper)
	helper.On("MakeHttpCall", mock.Anything, mock.Anything).Return(getResponse(getProto()), nil)
	_, err := service.GetMessageInInterval(context.Background(), &proto.GetMessageInIntervalRequest{})
	assert.NoError(t, err)
}

func Test_GetMessageInInterval_Error(t *testing.T) {
	helper := &mocks.Helper{}
	service := NewLogAnalyserService(helper)
	helper.On("MakeHttpCall", mock.Anything, mock.Anything).Return(nil, errors.New("Internal Server Error"))
	_, err := service.GetMessageInInterval(context.Background(), &proto.GetMessageInIntervalRequest{})
	assert.Error(t, err)
}

func Test_GetMessageInInterval_SuccessWithEmptyList(t *testing.T) {
	helper := &mocks.Helper{}
	service := NewLogAnalyserService(helper)
	resp := getProto()
	resp.Logs = []*proto.Log{}
	helper.On("MakeHttpCall", mock.Anything, mock.Anything).Return(getResponse(resp), nil)
	response, err := service.GetMessageInInterval(context.Background(), &proto.GetMessageInIntervalRequest{})
	assert.NoError(t, err)
	assert.True(t, response.Logs == nil)
}

func Test_GetMessageInInterval_NotFoundError(t *testing.T) {
	helper := &mocks.Helper{}
	service := NewLogAnalyserService(helper)
	helper.On("MakeHttpCall", mock.Anything, mock.Anything).Return(nil, errors.New("Not Found"))
	_, err := service.GetMessageInInterval(context.Background(), &proto.GetMessageInIntervalRequest{})
	assert.Error(t, err)
	assert.True(t, err.Error() == "Not Found")
}

func Test_GetMessageInInterval_SuccessWithValues(t *testing.T) {
	helper := &mocks.Helper{}
	service := NewLogAnalyserService(helper)
	helper.On("MakeHttpCall", mock.Anything, mock.Anything).Return(getResponse(getProto()), nil)
	resp, err := service.GetMessageInInterval(context.Background(), &proto.GetMessageInIntervalRequest{})
	assert.NoError(t, err)
	assert.True(t, len(resp.Logs) > 0)
}

func getResponse(resp proto.GetMessageInIntervalResponse) *http.Response {
	str, _ := json.Marshal(resp)
	return &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(str))),
	}
}

func getProto() proto.GetMessageInIntervalResponse {
	return proto.GetMessageInIntervalResponse{
		Date: "2022-10-06",
		Logs: []*proto.Log{
			&proto.Log{
				MessageHash: "abcd",
				Time:        "00:29:21.017",
			},
		},
	}
}
