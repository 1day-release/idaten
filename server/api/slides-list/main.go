package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func slides_list(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	personID := "123456acsdvdw"
	personName := "mym@gmail.com"
	old := 25

	person := PersonResponse{
		PersonID:   personID,
		PersonName: personName,
		Old:        old,
	}
	jsonBytes, _ := json.Marshal(person)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

type PersonResponse struct {
	PersonID   string `json:"personID"`
	PersonName string `json:"personName"`
	Old        int    `json:"old"`
}

func main() {
	lambda.Start(slides_list)
}
