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
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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
	// スライドパスを取得する
	path := request.Path
	getSlideID := strings.Split(path, "/")[2]
	fmt.Println(getSlideID)

	if len(bearerAccessTokenSplit) == 1 || getSlideID == "" {
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

	result, err := svc.Query(&dynamodb.QueryInput{
		TableName: aws.String("idaten-slides"),
		ExpressionAttributeNames: map[string]*string{
			"#ID":        aws.String("slide_id"),
			"#MARKDOWN":  aws.String("markdown"),
			"#SHAREMODE": aws.String("share_mode"),
			"#CREATEDAT": aws.String("created_at"),
			"#UPDATEDAT": aws.String("updated_at"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(getSlideID),
			},
		},
		KeyConditionExpression: aws.String("#ID = :id"),
		ProjectionExpression:   aws.String("#ID, #MARKDOWN, #SHAREMODE, #CREATEDAT, #UPDATEDAT"),
	})
	fmt.Println(result.Items)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	var responseUserData ResponseUserData

	for _, i := range result.Items {
		item := UserData{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		jsonData := UserData{
			SlideID:   item.SlideID,
			Markdown:  item.Markdown,
			ShareMode: item.ShareMode,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}

		responseUserData = append(responseUserData, jsonData)
	}

	jsonBytes, _ := json.Marshal(responseUserData)
	jsonString := string(jsonBytes)
	if jsonString == "null" {
		jsonString = "[]"
	}

	return events.APIGatewayProxyResponse{
		Body:       jsonString,
		Headers:    responseHeader,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type UserData struct {
	SlideID   string `json:"slide_id"`
	Markdown  string `json:"markdown"`
	ShareMode int    `json:"share_mode"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ResponseUserData []UserData
