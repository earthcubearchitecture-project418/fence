package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"../../internal/fence"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yterajima/go-sitemap"
)

var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)

type FenceEvent struct {
	Sitemap string `json:"sitemap"`
	Date    string `json:"date"`
}

func show(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the `isbn` query string parameter from the request and
	// validate it.
	// isbn := req.QueryStringParameters["isbn"]
	// if !isbnRegexp.MatchString(isbn) {
	// 	return clientError(http.StatusBadRequest)
	// }

	data := fence.MapData{}

	smapurl := req.QueryStringParameters["sitemap"]
	fd := req.QueryStringParameters["date"]

	smap, err := sitemap.Get(smapurl, nil)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Body:       "",
		}, err
	}

	c, o, err := fence.DateCheck(smap, fd)
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

// Add a helper for handling errors. This logs any error to os.Stderr
// and returns a 500 Internal Server Error response that the AWS API
// Gateway understands.
func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

// Similarly add a helper for send responses relating to client errors.
func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(show)
}
