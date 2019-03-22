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
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	goLambda "github.com/aws/aws-sdk-go/service/lambda"
)

func athorizationIdatenUser(accessToken UserAccessToken) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	client := goLambda.New(sess, &aws.Config{Region: aws.String("ap-northeast-1")})

	payload, err := json.Marshal(accessToken)
	//payload, err := json.Marshal(accessToken)
	if err != nil {
		fmt.Println("Error marshalling idaten-api-users-authorization request")
		os.Exit(0)
	}

	response, err := client.Invoke(&goLambda.InvokeInput{FunctionName: aws.String("idaten-api-users-authorization"), Payload: payload})
	if err != nil {
		fmt.Println("Error calling idaten-api-users-authorization")
		os.Exit(0)
	}

	var lambdaResponse LambdaResponse
	err = json.Unmarshal(response.Payload, &lambdaResponse)
	if err != nil {
		fmt.Println("Error unmarshalling idaten-api-users-authorization response")
		os.Exit(0)
	}

	fmt.Printf(lambdaResponse.Message)
}

func slidesList(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// CORSレスポンスヘッダを設定
	responseHeader := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
		"Content-Type":                 "application/json",
	}

	// AccessTokenを取得する
	bearerAccessToken := request.Headers["Authorization"]
	bearerAccessTokenSplit := strings.Split(bearerAccessToken, " ")
	// メールアドレスを取得
	requestEmail := request.QueryStringParameters["email"]

	if len(bearerAccessTokenSplit) == 1 || len(requestEmail) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Bad Request"}`,
			Headers:    responseHeader,
			StatusCode: 400,
		}, nil
	}

	accessToken := bearerAccessTokenSplit[1]
	userAccessToken := UserAccessToken{AccessToken: accessToken}
	athorizationIdatenUser(userAccessToken)

	// DynamoDBのテーブルと接続
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		fmt.Println("Got error creating session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create DynamoDB client
	svc := dynamodb.New(session)

	filter := expression.Name("email").Equal(expression.Value(requestEmail))
	project := expression.NamesList(expression.Name("slide_id"), expression.Name("email"), expression.Name("share_mode"), expression.Name("created_at"), expression.Name("updated_at"))
	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(project).Build()

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("idaten-slides"),
	}

	results, err := svc.Scan(params)
	fmt.Print(err)

	var responseUserData ResponseUserData

	for _, i := range results.Items {
		item := UserData{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		jsonData := UserData{
			Email:     item.Email,
			SlideID:   item.SlideID,
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

type UserData struct {
	SlideID   string `json:"slide_id"`
	Email     string `json:"email"`
	ShareMode int    `json:"share_mode"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ResponseUserData []UserData

type UserAccessToken struct {
	AccessToken string `json:"access_token"`
}

type LambdaResponse struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func main() {
	lambda.Start(slidesList)
}
