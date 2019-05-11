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
	// idatenUserテーブルからユーザ情報を取得
	lambdaResponse := athorizationIdatenUser(userAccessToken)

	//スライドIDを取得
	path := request.Path
	getSlideID := strings.Split(path, "/")[2]

	// lambdaresponseのメッセージからIdatenユーザ情報を取り出す
	idatenUserInfo := new(IdatenUserInfo)
	err := json.Unmarshal([]byte(lambdaResponse.Message), idatenUserInfo)

	if len(bearerAccessTokenSplit) == 1 || err != nil || getSlideID == "" {
		fmt.Println("get error:", err)
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Forbidden"}`,
			Headers:    responseHeader,
			StatusCode: 403,
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

	if len(result.Items) == 0 {
		fmt.Println("Not find ID")
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Not Found"}`,
			Headers:    responseHeader,
			StatusCode: 404,
		}, nil
	}

	_, err = svc.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"slide_id": {
				S: aws.String(getSlideID),
			},
			"email": {
				S: aws.String(idatenUserInfo.Email),
			},
		},
		TableName: aws.String("idaten-slides"),
	})
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
		Body:       `{"slide_id": "` + getSlideID + `"}`,
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
type UserAccessToken struct {
	AccessToken string `json:"access_token"`
}
type LambdaResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
