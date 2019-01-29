package main

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var path = request.Path
	var getSlideID = strings.Split(path, "/")[2]

	slideID := getSlideID
	markdown := "# Test\n\n## Test2"
	shareMode := 0
	createAt := "2019-01-01 00:00:00"
	updateAt := "2019-01-02 00:00:00"

	records := DynamoDBResponse{
		SlideID:   slideID,
		Markdown:  markdown,
		ShareMode: shareMode,
		CreateAt:  createAt,
		UpdateAt:  updateAt,
	}
	jsonBytes, _ := json.Marshal(records)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type DynamoDBResponse struct {
	SlideID   string `json:"SlideID"`
	Markdown  string `json:"Markdown"`
	ShareMode int    `json:"ShareMode"`
	CreateAt  string `json:"CreateAt"`
	UpdateAt  string `json:"UpdateAt"`
}
