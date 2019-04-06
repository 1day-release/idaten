package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

func scanUserTable(accessToken string) string {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		fmt.Println("Got error creating session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// DynamoDBのクライアントを作成
	svc := dynamodb.New(session)

	filt := expression.Name("access_token").Equal(expression.Value(accessToken))
	proj := expression.NamesList(expression.Name("user_id"), expression.Name("email"), expression.Name("access_token"), expression.Name("family_name"), expression.Name("given_name"), expression.Name("locale"), expression.Name("name"), expression.Name("picture"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("idaten-users"),
	}
	response, err := svc.Scan(params)

	idatenUserInfoData := IdatenUserInfoData{}
	for _, i := range response.Items {
		item := IdatenUserInfo{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		jsonData := IdatenUserInfo{
			ID:          item.ID,
			Email:       item.Email,
			AccessToken: item.AccessToken,
			Name:        item.Name,
			GivenName:   item.GivenName,
			FamilyName:  item.FamilyName,
			Picture:     item.Picture,
			Locale:      item.Locale,
		}

		idatenUserInfoData = append(idatenUserInfoData, jsonData)
	}

	jsonString := ""
	if len(idatenUserInfoData) == 0 {
		jsonString = "{}"
	} else {
		jsonBytes, _ := json.Marshal(idatenUserInfoData[0])
		jsonString = string(jsonBytes)
	}
	return jsonString
}

func deleteUserAccessToken(idatenUserInfo IdatenUserInfo) {
	idatenUserInfo.AccessToken = "None"

	av, err := dynamodbattribute.MarshalMap(idatenUserInfo)
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
		TableName: aws.String("idaten-users"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
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
	if len(bearerAccessTokenSplit) == 1 {
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Bad Request"}`,
			Headers:    responseHeader,
			StatusCode: 400,
		}, nil
	}
	accessToken := bearerAccessTokenSplit[1]

	// AccessTokenが含まれているレコードを取得
	jsonString := scanUserTable(accessToken)
	if jsonString == "{}" {
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Not Found"}`,
			Headers:    responseHeader,
			StatusCode: 404,
		}, nil
	}
	idatenUserInfo := IdatenUserInfo{}
	json.Unmarshal([]byte(jsonString), &idatenUserInfo)

	// AccessTokenのカラムのみを削除してレコードを更新
	deleteUserAccessToken(idatenUserInfo)

	return events.APIGatewayProxyResponse{
		Body:       `{"status": "OK"}`,
		Headers:    responseHeader,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type IdatenUserInfo struct {
	ID          string `json:"user_id"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
	GivenName   string `json:"given_name"`
	FamilyName  string `json:"family_name"`
	Picture     string `json:"picture"`
	Locale      string `json:"locale"`
}

type IdatenUserInfoData []IdatenUserInfo
