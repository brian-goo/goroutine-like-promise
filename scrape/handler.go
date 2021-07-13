package scrape

import (
	"context"

	"github.com/brian-goo/goroutine-like-promise/scrape/lambdaio"
	"github.com/brian-goo/goroutine-like-promise/scrape/scraper"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lambdaio.Init(&e)

	urls, err := lambdaio.GetPostData(&e)
	if err != "" {
		return lambdaio.GetErrorResponse(err)
	}

	res := make([]interface{}, len(urls))
	for k, v := range urls {
		res[k] = <-scraper.GetImg(&v)
	}

	return lambdaio.GetResponse(&res)
}
