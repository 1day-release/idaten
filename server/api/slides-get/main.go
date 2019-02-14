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
	var path = request.Path
	var getSlideID = strings.Split(path, "/")[2]
	fmt.Println(getSlideID)

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
		TableName: aws.String("idaten"),
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

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
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
