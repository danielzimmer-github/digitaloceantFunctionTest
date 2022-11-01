package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "github.com/lib/pq"
)

type CoolThings struct {
	Name string `json:"name"`
	Text string `json:"string"`
}

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	StatusCode int				`json:"statusCode,omitempty"`
	Headers map[string]string	`json:"headers,omitempty"`
	Body string 				`json:"body,omitempty"`
}



func Main(in CoolThings) (*Response, error) {
	if in.Name == "" {
		in.Name = "stranger"
	}
	if in.Text == "" {
		return &Response{
			Body: fmt.Sprintf("No text submitted"),
		}, nil
	}



	return &Response{
		Body: fmt.Sprintf("Authenticate %s", in.Name),
	}, nil
}