package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(userAccessToken UserAccessToken) (Response, error) {
	fmt.Println(userAccessToken.AccessToken)
	return Response{
		Message: userAccessToken.AccessToken,
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type UserAccessToken struct {
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
	Ok      bool   `json:"ok"`
}
