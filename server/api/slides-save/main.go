package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

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

func getNowTime() string {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	nowDateTime := time.Now().In(jst)
	const dateFormat = "2006-01-02 15:04:05"
	nowDateTimeString := nowDateTime.Format(dateFormat)

	return nowDateTimeString
}

// FindGetItem ... find item include exist slide id
func FindGetItem(svc *dynamodb.DynamoDB, getSlideID string) (int, string) {
	status := 200
	jsonString := "ok"

	fmt.Println(getSlideID)
	input := &dynamodb.GetItemInput{
		TableName: aws.String("idaten-slides"),
		Key: map[string]*dynamodb.AttributeValue{
			"slide_id": {
				S: aws.String(getSlideID),
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
func UpdateItemInput(svc *dynamodb.DynamoDB, getSlideID string, MarkDown string, ShareMode string, nowDateTime string) (int, string) {
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
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set #markdown = :markdown, #share_mode = :share_mode, #updated_at = :updated_at"),
	}

	_, err := svc.UpdateItem(input)

	if err != nil {
		return 400, `{"status": "Bad Request"}`
	}
	return 200, `{"slide_id": "` + getSlideID + `","updated_at":"` + nowDateTime + `"}`
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

	// Bodyの内容を取得
	jsonBytes := ([]byte)(request.Body)
	requestData := new(RequestData)
	err := json.Unmarshal(jsonBytes, requestData)
	//スライドIDを取得
	getSlideID := strings.Split(request.Path, "/")[2]

	// lambdaresponseのコードが403の場合にはレスポンスを返す
	if lambdaResponse.Code != 200 && (requestData.ShareMode == 0 || requestData.ShareMode == 1) {
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Forbidden"}`,
			Headers:    responseHeader,
			StatusCode: lambdaResponse.Code,
		}, nil
	}

	// 現在時刻を取得
	nowDateTime := getNowTime()

	// DynamoDBのセッションを作成
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	// Create DynamoDB client
	svc := dynamodb.New(sess)

	if err != nil {
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Bad Request"}`,
			Headers:    responseHeader,
			StatusCode: 400,
		}, nil
	}

	// share_modeが2（編集可能）である場合に更新する
	if lambdaResponse.Code != 200 && requestData.ShareMode == 2 {
		// find exist slide_id item
		status, jsonString := FindGetItem(svc, getSlideID)
		if status != 200 {
			return events.APIGatewayProxyResponse{
				Body:       jsonString,
				Headers:    responseHeader,
				StatusCode: status,
			}, nil
		}

		// update Item
		status, jsonString = UpdateItemInput(svc, getSlideID, requestData.MarkDown, strconv.Itoa(requestData.ShareMode), nowDateTime)

		return events.APIGatewayProxyResponse{
			Body:       jsonString,
			Headers:    responseHeader,
			StatusCode: status,
		}, nil
	}

	// lambdaresponseのメッセージからIdatenユーザ情報を取り出す
	idatenUserInfo := new(IdatenUserInfo)
	responseErr := json.Unmarshal([]byte(lambdaResponse.Message), idatenUserInfo)
	fmt.Println("Bearer Token length: " + strconv.Itoa(len(bearerAccessTokenSplit)))
	if len(bearerAccessTokenSplit) == 1 || err != nil || responseErr != nil || getSlideID == "" {
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Bad Request"}`,
			Headers:    responseHeader,
			StatusCode: 400,
		}, nil
	}

	// find exist slide_id item
	status, jsonString := FindGetItem(svc, getSlideID)
	if status != 200 {
		return events.APIGatewayProxyResponse{
			Body:       jsonString,
			Headers:    responseHeader,
			StatusCode: status,
		}, nil
	}

	// update Item
	status, jsonString = UpdateItemInput(svc, getSlideID, requestData.MarkDown, strconv.Itoa(requestData.ShareMode), nowDateTime)

	return events.APIGatewayProxyResponse{
		Body:       jsonString,
		Headers:    responseHeader,
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
