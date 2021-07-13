## Goroutine like Promise.all of javascript

- returns an slice of ordered results of scraped URLs of images
- can be deployed as microservice of aws lambda behind aws api gateway

```bash
// to test locally
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/brian-goo/goroutine-like-promise/scrape"
	"github.com/brian-goo/goroutine-like-promise/scrape/lambdaio"
)

func main() {
	var c context.Context
	post := `
		{"urls": ["https://www.yomiuri.co.jp/national/20210628-OYT1T50097/", "http://www.asahi.com/articles/ASP6X4640P6XUNHB001.html"]}
	`

	res, _ := scrape.Handler(c, events.APIGatewayProxyRequest{Body: post})
	js, _ := lambdaio.Encode(res)
	fmt.Printf("%s\n", js)
}
```

```bash
// to use with aws lambda main func
package main

import (
	"github.com/brian-goo/goroutine-like-promise/scrape"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(scrape.Handler)
}
```