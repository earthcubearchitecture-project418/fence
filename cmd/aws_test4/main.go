package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"../../internal/fence"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yterajima/go-sitemap"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	HTTPMethodNotSupported = errors.New("no name was provided in the HTTP body")
)

// HandleRequest is the test handler for now
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	data := fence.MapData{}

	fmt.Printf("Body size = %d. \n", len(request.Body))
	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("  %s: %s\n", key, value)
	}
	fmt.Printf("Path: %s \n", request.Path)

	if request.HTTPMethod == "GET" {
		fmt.Printf("GET METHOD\n")

		smapurl := request.QueryStringParameters["sitemap"]
		fd := request.QueryStringParameters["date"]

		fmt.Printf("-- POST METHOD\n")
		smap, err := sitemap.Get(smapurl, nil)
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       "",
			}, err
		}

		c, o, err := fence.DateCheck(smap, fd)
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       "",
			}, err
		}

		data = fence.MapData{MapLen: len(smap.URL), CheckLen: len(c), ErrorLen: len(o), URLs: c, ErrorURLs: o}

		j, err := json.MarshalIndent(data, " ", "")
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       "",
			}, err
		}

		return events.APIGatewayProxyResponse{Body: string(j), StatusCode: 200}, nil
	} else if request.HTTPMethod == "POST" {

		smapurl := request.QueryStringParameters["sitemap"]
		fd := request.QueryStringParameters["date"]

		fmt.Printf("-- POST METHOD\n")
		smap, err := sitemap.Get(smapurl, nil)
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       "",
			}, err
		}

		c, o, err := fence.DateCheck(smap, fd)
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       "",
			}, err
		}

		data = fence.MapData{MapLen: len(smap.URL), CheckLen: len(c), ErrorLen: len(o), URLs: c, ErrorURLs: o}

		_, err = json.Marshal(data)
		if err != nil {
			fmt.Println(err)
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusOK,
				Body:       "",
			}, err
		}

		return events.APIGatewayProxyResponse{Body: "why does this work?", StatusCode: 200}, nil
	} else {
		fmt.Printf("NEITHER\n")
		return events.APIGatewayProxyResponse{}, HTTPMethodNotSupported
	}
}

func main() {
	lambda.Start(HandleRequest)
}
