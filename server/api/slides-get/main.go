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
	goLambda "github.com/aws/aws-sdk-go/service/lambda"
)

func athorizationIdatenUser(accessToken UserAccessToken) (lambdaResponse LambdaResponse) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := goLambda.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})

	payload, err := json.Marshal(accessToken)
	if err != nil {
		fmt.Println("Error marshalling idaten-api-users-authorization request")
		os.Exit(0)
	}

	response, err := client.Invoke(&goLambda.InvokeInput{FunctionName: aws.String("idaten-api-users-authorization"), Payload: payload})
	if err != nil {
		fmt.Println("Error calling idaten-api-users-authorization")
		os.Exit(0)
	}

	//var lambdaResponse LambdaResponse
	err = json.Unmarshal(response.Payload, &lambdaResponse)
	if err != nil {
		fmt.Println("Error unmarshalling idaten-api-users-authorization response")
		os.Exit(0)
	}

	return lambdaResponse
}

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
	accessToken := bearerAccessTokenSplit[1]
	userAccessToken := UserAccessToken{AccessToken: accessToken}

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

	// idatenUserテーブルからユーザ情報を取得
	lambdaResponse := athorizationIdatenUser(userAccessToken)

	// lambdaresponseのコードが403の場合にはレスポンスを返す
	if lambdaResponse.Code != 200 {
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Forbidden"}`,
			Headers:    responseHeader,
			StatusCode: lambdaResponse.Code,
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
			"#COVER":     aws.String("cover"),
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
		ProjectionExpression:   aws.String("#ID, #COVER, #MARKDOWN, #SHAREMODE, #CREATEDAT, #UPDATEDAT"),
	})
	fmt.Println(result.Items)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	var jsonData UserData
	var ResponseStatusCode int
	var jsonString string

	if len(result.Items) == 0 {
		ResponseStatusCode = 404
		jsonString = `{"status": "Not Found"}`
	} else {
		for _, i := range result.Items {
			ResponseStatusCode = 200
			item := UserData{}

			err = dynamodbattribute.UnmarshalMap(i, &item)

			if err != nil {
				fmt.Println("Got error unmarshalling:")
				fmt.Println(err.Error())
				os.Exit(1)
			}

			jsonData = UserData{
				SlideID:   item.SlideID,
				Cover:     item.Cover,
				Markdown:  item.Markdown,
				ShareMode: item.ShareMode,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			}

			jsonBytes, _ := json.Marshal(jsonData)
			jsonString = string(jsonBytes)
			break
		}
	}

	return events.APIGatewayProxyResponse{
		Body:       jsonString,
		Headers:    responseHeader,
		StatusCode: ResponseStatusCode,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type UserData struct {
	SlideID   string `json:"slide_id"`
	Cover     string `json:"cover"`
	Markdown  string `json:"markdown"`
	ShareMode int    `json:"share_mode"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserAccessToken struct {
	AccessToken string `json:"access_token"`
}

type LambdaResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
