package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

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

	// Bodyの内容を取得する
	jsonBytes := ([]byte)(request.Body)
	requestData := new(RequestData)

	if err := json.Unmarshal(jsonBytes, requestData); err != nil || len(bearerAccessTokenSplit) == 1 {
		fmt.Println("JSON Unmarshal error:", err)
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

	// lambdaresponseのメッセージからIdatenユーザ情報を取り出す
	idatenUserInfo := new(IdatenUserInfo)
	err := json.Unmarshal([]byte(lambdaResponse.Message), idatenUserInfo)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Bad Request"}`,
			Headers:    responseHeader,
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
		Email:     idatenUserInfo.Email,
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
		TableName: aws.String("idaten-slides"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	jsonString := `{"slide_id":"` + slideID + `"}`
	return events.APIGatewayProxyResponse{
		Body:       jsonString,
		Headers:    responseHeader,
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

type UserAccessToken struct {
	AccessToken string `json:"access_token"`
}
type LambdaResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
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
