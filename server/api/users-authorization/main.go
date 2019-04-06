package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func searchIdatenUser(accessToken string) (jsonString string, statuscode int) {
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

	// Queryの作成
	result, err := svc.Query(&dynamodb.QueryInput{
		TableName: aws.String("idaten-users"),
		ExpressionAttributeNames: map[string]*string{
			"#USERID":      aws.String("user_id"),
			"#EMAIL":       aws.String("email"),
			"#ACCESSTOKEN": aws.String("access_token"),
			"#FAMILYNAME":  aws.String("family_name"),
			"#GIVENNAME":   aws.String("given_name"),
			"#LOCALE":      aws.String("locale"),
			"#NAME":        aws.String("name"),
			"#PICTURE":     aws.String("picture"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":access_token": {
				S: aws.String(accessToken),
			},
		},
		KeyConditionExpression: aws.String("#ACCESSTOKEN = :access_token"),
		// グローバルセカンダリインデックス
		IndexName:            aws.String("access_token-index"),
		ProjectionExpression: aws.String("#USERID, #EMAIL, #ACCESSTOKEN, #FAMILYNAME, #GIVENNAME, #LOCALE, #NAME, #PICTURE"),
	})

	// 取得したデータからstructに変換
	idatenUserInfo := IdatenUserInfo{}
	for _, i := range result.Items {
		err = dynamodbattribute.UnmarshalMap(i, &idatenUserInfo)
		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	// jsonに変換
	jsonBytes, _ := json.Marshal(idatenUserInfo)
	jsonString = string(jsonBytes)
	fmt.Printf(jsonString)

	// データが無い場合には，空配列としてステータスコードを403とする
	statuscode = 200
	if len(idatenUserInfo.ID) == 0 {
		jsonString = "[]"
		statuscode = 403
	}
	return jsonString, statuscode
}

func handler(requestUserInfo RequestUserInfo) (Response, error) {
	fmt.Println(requestUserInfo.AccessToken)

	idatenUserInfo, statuscode := searchIdatenUser(requestUserInfo.AccessToken)
	return Response{
		Message: idatenUserInfo,
		Code:    statuscode,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type RequestUserInfo struct {
	AccessToken string `json:"access_token"`
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

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
