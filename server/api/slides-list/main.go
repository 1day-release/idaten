package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

func slidesList(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// CORSレスポンスヘッダを設定
	responseHeader := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
		"Content-Type":                 "application/json",
	}

	// AccessTokenを取得する
	bearerAccessToken := request.Headers["Authorization"]

	// アクセストークンが無い場合には403とする
	if len(bearerAccessToken) == 0 {
		return events.APIGatewayProxyResponse{
			Body:       `{"status": "Forbidden"}`,
			Headers:    responseHeader,
			StatusCode: 403,
		}, nil
	}

	bearerAccessTokenSplit := strings.Split(bearerAccessToken, " ")
	accessToken := bearerAccessTokenSplit[1]
	userAccessToken := UserAccessToken{AccessToken: accessToken}
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

	filter := expression.Name("email").Equal(expression.Value(idatenUserInfo.Email))
	project := expression.NamesList(expression.Name("slide_id"), expression.Name("email"), expression.Name("cover"), expression.Name("share_mode"), expression.Name("created_at"), expression.Name("updated_at"))
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
			Cover:     item.Cover,
			SlideID:   item.SlideID,
			ShareMode: item.ShareMode,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}

		responseUserData = append(responseUserData, jsonData)
	}

	// created_atを降順にソートする
	sort.SliceStable(responseUserData, func(i, j int) bool { return responseUserData[i].CreatedAt > responseUserData[j].CreatedAt })
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
type UserData struct {
	SlideID   string `json:"slide_id"`
	Email     string `json:"email"`
	Cover     string `json:"cover"`
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
	Code    int    `json:"code"`
}

func main() {
	lambda.Start(slidesList)
}
