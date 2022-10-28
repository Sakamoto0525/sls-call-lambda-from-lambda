package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	sdklambda "github.com/aws/aws-sdk-go/service/lambda"
	"github.com/kelseyhightower/envconfig"
)

type Response events.APIGatewayProxyResponse

type Config struct {
	WorldURL string `envconfig:"WORLD_URL"`
}

func Handler(ctx context.Context) (Response, error) {
	fmt.Print("====== Start hello functions ======")

	var buf bytes.Buffer

	config := &Config{}
	envconfig.MustProcess("", config)
	fmt.Printf("config: %s", config)

	// sess := aws.NewSession(&aws.Config{
	// 	Region: aws.String("ap-northeast-1"),
	// })
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
		// Profile: "default",
	}))
	lambda := sdklambda.New(sess)

	input := &sdklambda.InvokeInput{
		FunctionName: aws.String(config.WorldURL),
	}

	res, err := lambda.Invoke(input)
	if err != nil {
		fmt.Errorf("lambda.Invoke Error: %s", err)
	}
	fmt.Printf("lambda.Invoke Succeded: %s", res)

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	fmt.Print("====== end hello functions ======")

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
