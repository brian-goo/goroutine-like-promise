package lambdaio

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func Encode(i interface{}) ([]byte, error) {
	data, err := json.Marshal(i)
	if err != nil {
		return []byte(`{"status": "json encode failed"}`), err
	}
	return data, nil
}

func Init(i interface{}) {
	js, err := Encode(i)
	if err != nil {
		fmt.Printf("json encode for printing request failed: %s\n", err)
		return
	}
	fmt.Printf("event: %s\n", js)
}

func GetPostData(e *events.APIGatewayProxyRequest) ([]string, string) {
	var m map[string][]string
	err := json.Unmarshal([]byte(e.Body), &m)
	if err != nil {
		fmt.Printf("json decode for post data failed: %s\n", err)
		return []string{""}, "json decode for post data failed"
	}

	urls, ok := m["urls"]
	if !ok {
		fmt.Println("no key of urls in post data")
		return urls, "no key of urls in post data"
	}

	return urls, ""
}

func GetResponse(body interface{}) (events.APIGatewayProxyResponse, error) {
	status := 200
	js, err := Encode(body)
	if err != nil {
		status = 500
		fmt.Printf("json encode for response failed: %s\n", err)
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
		IsBase64Encoded: false,
		StatusCode:      status,
		Body:            string(js),
	}, nil
}

func GetErrorResponse(e string) (events.APIGatewayProxyResponse, error) {
	status := 500
	m := map[string]string{
		"status": e,
	}

	js, err := Encode(m)
	if err != nil {
		fmt.Printf("json encode for response failed: %s\n", err)
	}

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
			"Content-Type":                "application/json",
		},
		IsBase64Encoded: false,
		StatusCode:      status,
		Body:            string(js),
	}, nil
}
