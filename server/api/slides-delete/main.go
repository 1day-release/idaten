package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// CORSレスポンスヘッダを設定
	responseHeader := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
		"Content-Type":                 "application/json",
	}

	// AccessTokenを取得する
	bearerAccessToken := request.Headers["Authorization"]
	bearerAccessTokenSplit := strings.Split(bearerAccessToken, " ")
	// Bodyの内容を取得
	jsonBytes := ([]byte)(request.Body)
	requestData := new(RequestData)
	err := json.Unmarshal(jsonBytes, requestData)
	//スライドIDを取得
	path := request.Path
	getSlideID := strings.Split(path, "/")[2]

	if len(bearerAccessTokenSplit) == 1 || err != nil || getSlideID == "" {
		fmt.Println("get error:", err)
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Bad Request"}`,
			Headers:    responseHeader,
			StatusCode: 400,
		}, nil
	}

	// DynamoDBのテーブルと接続
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		fmt.Println("Got error creating session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	svc := dynamodb.New(session)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"slide_id": {
				S: aws.String(getSlideID),
			},
			"email": {
				S: aws.String(requestData.Email),
			},
		},
		TableName: aws.String("idaten-slides"),
	}

	_, err = svc.DeleteItem(input)
	fmt.Println(err)

	// データが無いときでも200を返してしまう
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Not Found"}`,
			Headers:    responseHeader,
			StatusCode: 404,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       `{"status": "200"}`,
		Headers:    responseHeader,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type RequestData struct {
	Email string `json:"email"`
}
