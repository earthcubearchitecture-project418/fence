package main

import (
	"encoding/json"
	"net/http"

	"../../internal/fence"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yterajima/go-sitemap"
)

type FenceEvent struct {
	Sitemap string `json:"sitemap"`
	Date    string `json:"date"`
}

// func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func HandleLambdaEvent(event FenceEvent) (events.APIGatewayProxyResponse, error) {
	data := fence.MapData{}

	smap, err := sitemap.Get(event.Sitemap, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "",
		}, err
	}

	c, o, err := fence.DateCheck(smap, event.Date)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "",
		}, err
	}

	data = fence.MapData{MapLen: len(smap.URL), CheckLen: len(c), ErrorLen: len(o), URLs: c, ErrorURLs: o}

	j, err := json.Marshal(data)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(j),
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
