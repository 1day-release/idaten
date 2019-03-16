package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func connectToGoogle(requestGoogleData RequestGoogleClientData) ResponseGoogleServerData {
	reflectType, reflectValue := reflect.TypeOf(requestGoogleData), reflect.ValueOf(requestGoogleData)
	jsonBody := "{"
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		jsonKey := field.Tag.Get("json")
		jsonValue := reflectValue.Field(i).Interface()
		jsonElement := "\"" + jsonKey + "\": " + "\"" + jsonValue.(string) + "\","
		if i == reflectType.NumField()-1 {
			jsonElement = strings.Replace(jsonElement, ",", "", 1)
		}

		jsonBody += jsonElement
	}
	jsonBody += "}"

	const URL = "https://accounts.google.com/o/oauth2/token"
	request, err := http.NewRequest(
		"POST",
		URL,
		bytes.NewBuffer([]byte(jsonBody)),
	)
	if err != nil {
	}
	fmt.Printf(jsonBody)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	responseGoogleServerData := ResponseGoogleServerData{}
	if err = json.NewDecoder(response.Body).Decode(&responseGoogleServerData); err != nil {
	}
	defer response.Body.Close()

	return responseGoogleServerData
}

func getUserInfoFromGoogle(accessToken string) ResponseGoogleUserInfo {
	RequestBearer := "Bearer " + accessToken
	const URL = "https://www.googleapis.com/oauth2/v1/userinfo"
	request, err := http.NewRequest(
		"GET",
		URL,
		nil,
	)
	request.Header.Set("Authorization", RequestBearer)
	if err != nil {
	}

	client := &http.Client{}
	response, err := client.Do(request)

	responseGoogleUserInfo := ResponseGoogleUserInfo{}
	if err = json.NewDecoder(response.Body).Decode(&responseGoogleUserInfo); err != nil {
	}

	fmt.Print(responseGoogleUserInfo)
	return responseGoogleUserInfo
}

func UserInfoItemInput(responseGoogleServerData ResponseGoogleServerData, responseGoogleUserInfo ResponseGoogleUserInfo) {
	idatenUserInfo := IdatenUserInfo{
		ID:          responseGoogleUserInfo.ID,
		Email:       responseGoogleUserInfo.Email,
		AccessToken: responseGoogleServerData.AccessToken,
		Name:        responseGoogleUserInfo.Name,
		GivenName:   responseGoogleUserInfo.GivenName,
		FamilyName:  responseGoogleUserInfo.FamilyName,
		Picture:     responseGoogleUserInfo.Picture,
		Locale:      responseGoogleUserInfo.Locale,
	}

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

	jsonBytes := ([]byte)(request.Body)
	requestClientData := new(RequestClientData)

	if err := json.Unmarshal(jsonBytes, requestClientData); err != nil {
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

	// Googleサーバに問い合わせるための要素
	requestGoogleData := &RequestGoogleClientData{
		Code:         requestClientData.Code,
		ClientID:     os.Getenv("ClientID"),
		ClientSecret: os.Getenv("ClientSecret"),
		RedirectURL:  "https://api.idaten.1day-release.com/auth/login",
		GrantType:    "authorization_code",
	}

	responseGoogleServerData := connectToGoogle(*requestGoogleData)

	// AccessTokenを用いてGoogleにユーザ情報の問い合わせ
	responseGoogleUserInfo := getUserInfoFromGoogle(responseGoogleServerData.AccessToken)
	jsonBytes, _ = json.Marshal(responseGoogleUserInfo)

	// DynamoDBにAccessTokenやユーザ情報を格納
	UserInfoItemInput(responseGoogleServerData, responseGoogleUserInfo)

	// AccessTokenを返却する
	return events.APIGatewayProxyResponse{
		Body: "{\"access_token\":\"" + responseGoogleServerData.AccessToken + "\"}",
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

type RequestClientData struct {
	Code string `json:"code"`
}

type RequestGoogleClientData struct {
	Code         string `json:"code"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_uri"`
	GrantType    string `json:"grant_type"`
}

type ResponseGoogleServerData struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	IDToken      string `json:"id_token"`
}

type ResponseGoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
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
