package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getNowTime() string {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	nowDateTime := time.Now().In(jst)
	const dateFormat = "2006-01-02 15:04:05"
	nowDateTimeString := nowDateTime.Format(dateFormat)

	return nowDateTimeString
}

func getSlideID(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	const (
		letterIdxBits = 6
		letterIdxMask = 1<<letterIdxBits - 1
		letterIdxMax  = 63 / letterIdxBits
	)

	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(request)
	fmt.Println(request.Body)
	fmt.Println(request.Headers)
	fmt.Println(request.IsBase64Encoded)
	fmt.Println(request.Path)
	fmt.Println(request.PathParameters)
	fmt.Println(request.QueryStringParameters)
	fmt.Println(request.RequestContext)
	fmt.Println(request.Resource)
	fmt.Println(request.StageVariables)

	jsonBytes := ([]byte)(request.Body)
	requestData := new(RequestData)

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
	// スライドIDを生成
	slideID := getSlideID(64)

	// create item dor dynamodb
	item := PutItem{
		SlideID:   slideID,
		Email:     requestData.Email,
		ShareMode: requestData.ShareMode,
		CreatedAt: nowDateTime,
	}
	av, err := dynamodbattribute.MarshalMap(item)

	session, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		fmt.Println("Got error creating session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	svc := dynamodb.New(session)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("idaten"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	jsonString := `{"slide_id":"` + slideID + `"}`
	return events.APIGatewayProxyResponse{
		Body: jsonString,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
			"Content-Type":                 "application/json",
		},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type RequestData struct {
	Email     string `json:"email"`
	ShareMode int    `json:"share_mode"`
}

type PutItem struct {
	SlideID   string `json:"slide_id"`
	Email     string `json:"email"`
	ShareMode int    `json:"share_mode"`
	CreatedAt string `json:"created_at"`
}

type ResponseSlideID struct {
	SlideID string `json:"slide_id"`
}
