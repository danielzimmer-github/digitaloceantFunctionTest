package main

import (
	"fmt"
)

type Request struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type Response struct {
	StatusCode 	int 				`json:"statusCode,omitempty"`
	Headers 	map[string]string 	`json:"headers,omitempty"`
	Body 		string 				`json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	if in.Name == "" {
		in.Name = "stranger"
	}

	return &Response{
		Body: fmt.Sprintf("Hello %s %s", in.Name, in.Text),
	}, nil
}