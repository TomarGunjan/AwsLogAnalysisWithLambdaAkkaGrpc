package helper

import (
	"context"
	"net/http"
)

// This file contains helper function for the service to make HTTP calls
type Helper interface {
	MakeHttpCall(ctx context.Context, url string) (*http.Response, error)
}

type helper struct {
}

func GetHelper() Helper {
	return &helper{}
}

func (s *helper) MakeHttpCall(ctx context.Context, url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
