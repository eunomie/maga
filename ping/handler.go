package main

import (
	"encoding/json"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

type LambdaEvent struct {
	Body    string      `json:"body"`
	Headers interface{} `json:"headers"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Query   interface{} `json:"query"`
}

type Output struct {
	Status string `json:"status"`
}

func handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	var event LambdaEvent
	err := json.Unmarshal(evt, &event)
	if err != nil {
		return nil, err
	}

	var out Output
	out.Status = "ok"

	return out, nil
}

func init() {
	runtime.HandleFunc(handle)
}

func main() {}
