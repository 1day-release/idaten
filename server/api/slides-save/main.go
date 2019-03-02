package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func getNowTime() string {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	nowDateTime := time.Now().In(jst)
	const dateFormat = "2006-01-02 15:04:05"
	nowDateTimeString := nowDateTime.Format(dateFormat)

	return nowDateTimeString
}

// FindGetItem ... find item include exist slide id
func FindGetItem(svc *dynamodb.DynamoDB, getSlideID string, Email string) (int, string) {
	status := 200
	jsonString := "ok"

	fmt.Println(getSlideID)
	fmt.Println(Email)
	input := &dynamodb.GetItemInput{
		TableName: aws.String("idaten-slides"),
		Key: map[string]*dynamodb.AttributeValue{
			"slide_id": {
				S: aws.String(getSlideID),
			},
			"email": {
				S: aws.String(Email),
			},
		},
	}

	result, err := svc.GetItem(input)

	if err != nil {
		fmt.Println("[GetItem Error]", err)
		status = 400
		jsonString = `{"status": "Bad Request"}`
	}
	if len(result.Item) == 0 {
		status = 404
		jsonString = `{"status": "Not Found"}`
	}
	return status, jsonString
}

// UpdateItemInput ... update dynamodb table
func UpdateItemInput(svc *dynamodb.DynamoDB, getSlideID string, Email string, MarkDown string, ShareMode string, nowDateTime string) (int, string) {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":markdown": {
				S: aws.String(MarkDown),
			},
			":share_mode": {
				N: aws.String(ShareMode),
			},
			":updated_at": {
				S: aws.String(nowDateTime),
			},
		},

		ExpressionAttributeNames: map[string]*string{
			"#markdown":   aws.String("markdown"),
			"#share_mode": aws.String("share_mode"),
			"#updated_at": aws.String("updated_at"),
		},

		TableName: aws.String("idaten-slides"),
		Key: map[string]*dynamodb.AttributeValue{
			"slide_id": {
				S: aws.String(getSlideID),
			},
			"email": {
				S: aws.String(Email),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set #markdown = :markdown, #share_mode = :share_mode, #updated_at = :updated_at"),
	}

	_, err := svc.UpdateItem(input)

	if err != nil {
		return 400, `{"status": "Bad Request"}`
	}
	return 200, `{"updated_at":"` + nowDateTime + `"}`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request.Body)
	jsonBytes := ([]byte)(request.Body)
	requestData := new(RequestData)

	getSlideID := strings.Split(request.Path, "/")[2]

	if err := json.Unmarshal(jsonBytes, requestData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return events.APIGatewayProxyResponse{
			Body: `{"status": "Bad Request"}`,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
				"Content-Type":                 "application/json",
			},
			StatusCode: 400,
		}, nil
	}

	// 現在時刻を取得
	nowDateTime := getNowTime()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	if err != nil {
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{
			Body: `{"status": "Bad Request"}`,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
				"Content-Type":                 "application/json",
			},
			StatusCode: 400,
		}, nil
	}

	// find exist slide_id item
	status, jsonString := FindGetItem(svc, getSlideID, requestData.Email)
	if status != 200 {
		return events.APIGatewayProxyResponse{
			Body: jsonString,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  "*",
				"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
				"Content-Type":                 "application/json",
			},
			StatusCode: status,
		}, nil
	}

	// update Item
	status, jsonString = UpdateItemInput(svc, getSlideID, requestData.Email, requestData.MarkDown, strconv.Itoa(requestData.ShareMode), nowDateTime)

	return events.APIGatewayProxyResponse{
		Body: jsonString,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
			"Content-Type":                 "application/json",
		},
		StatusCode: status,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type RequestData struct {
	getSlideID string
	Email      string `json:"email"`
	MarkDown   string `json:"markdown"`
	ShareMode  int    `json:"share_mode"`
}

type PutItem struct {
	SlideID   string `json:"slide_id"`
	Email     string `json:"email"`
	MarkDown  string `json:"markdown"`
	ShareMode int    `json:"share_mode"`
	UpdatedAt string `json:"updated_at"`
}

type ResponseUpdatedAt struct {
	UpdatedAt string `json:"updated_at"`
}
