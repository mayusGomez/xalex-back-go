package main

type Request struct {
	company    string
	pageNumber int
	pageSize   int
}

type Response struct {
	StatusCode string      `json:"CODE"`
	Body       interface{} `json:"DATA"`
	Message    string      `json:"DESC"`
}
