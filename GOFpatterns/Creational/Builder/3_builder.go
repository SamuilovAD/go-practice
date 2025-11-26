package main

import (
	"fmt"
	"sync"
)

func main() {
	builder := NewRequestBuilder()
	request := builder.
		SetMethod("POST").
		SetURL("https://api.example.com/data").
		AddHeader("Authorization", "Bearer token").
		AddHeader("Content-Type", "application/json").
		AddQueryParam("sort", "desc").
		SetBody(`{"name": "alex"}`).
		Build()

	fmt.Printf("Built request: %+v\n", request)
}

type Request struct {
	Method  string
	URL     string
	Headers map[string]string
	Query   map[string]string
	Body    string
}

type RequestBuilder struct {
	mutex   sync.Mutex
	request Request
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		request: Request{
			Headers: make(map[string]string),
			Query:   make(map[string]string),
		},
	}
}

func (builder *RequestBuilder) SetMethod(method string) *RequestBuilder {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()
	builder.request.Method = method
	return builder
}

func (builder *RequestBuilder) SetURL(url string) *RequestBuilder {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()
	builder.request.URL = url
	return builder
}

func (builder *RequestBuilder) AddHeader(key, value string) *RequestBuilder {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()
	builder.request.Headers[key] = value
	return builder
}

func (builder *RequestBuilder) AddQueryParam(key, value string) *RequestBuilder {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()
	builder.request.Query[key] = value
	return builder
}

func (builder *RequestBuilder) SetBody(body string) *RequestBuilder {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()
	builder.request.Body = body
	return builder
}

func (builder *RequestBuilder) Build() Request {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()
	requestCopy := builder.request
	return requestCopy
}
