package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fale/swlcs/objects"
	"github.com/fale/swlcs/strategies"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	lambda.Start(HandleRequest)
}

type CommentRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Body     string `json:"body"`
	Resource string `json:"resource"`
}

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "POST":
		res, err := ProcessRequest(ctx, req)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 200,
				Headers: map[string]string{
					"Access-Control-Allow-Origin":  os.Getenv("ACAO"),
					"Content-Type":                 "application/json",
					"Access-Control-Allow-Methods": "POST, OPTIONS",
					"Access-Control-Allow-Headers": "content-type",
				},
				Body: "{'error': '" + err.Error() + "'}",
			}, nil
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  os.Getenv("ACAO"),
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Methods": "POST, OPTIONS",
				"Access-Control-Allow-Headers": "content-type",
			},
			Body: "{'result': '" + res + "'}",
		}, nil
	case "OPTIONS":
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  os.Getenv("ACAO"),
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Methods": "POST, OPTIONS",
				"Access-Control-Allow-Headers": "content-type",
			},
		}, nil
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin":  os.Getenv("ACAO"),
				"Content-Type":                 "application/json",
				"Access-Control-Allow-Methods": "POST, OPTIONS",
				"Access-Control-Allow-Headers": "content-type",
			},
			Body: "{'error': 'method " + req.HTTPMethod + " not supported'}",
		}, nil
	}
}

func ProcessRequest(ctx context.Context, req events.APIGatewayProxyRequest) (string, error) {
	payload := CommentRequest{}
	if err := json.Unmarshal([]byte(req.Body), &payload); err != nil {
		return fmt.Sprintf("impossible to decode JSON: %s", err), err
	}
	tz, err := time.LoadLocation("UTC")
	if err != nil {
		return fmt.Sprintf("impossible to get the correct timezone informations: %s", err), err
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")})
	tc := oauth2.NewClient(ctx, ts)

	st, err := strategies.Init(
		os.Getenv("GIT_STRATEGY"),
		ctx,
		&objects.Repository{
			GitHubClient: github.NewClient(tc),
			Owner:        os.Getenv("GITHUB_REPOSITORY_OWNER"),
			Name:         os.Getenv("GITHUB_REPOSITORY_NAME"),
			Branch:       os.Getenv("GITHUB_REPOSITORY_BRANCH"),
		},
		&objects.Comment{
			Resource:    payload.Resource,
			AuthorName:  payload.Name,
			AuthorEmail: payload.Email,
			Body:        strings.ReplaceAll(payload.Body, `\n`, "\n"),
			Time:        time.Now().In(tz),
		},
	)
	if err != nil {
		return fmt.Sprintf("an error occurred while initializing the strategy implementation: %s", err), err
	}
	if err := st.Execute(); err != nil {
		return fmt.Sprintf("an error occurred while executing git commands: %s", err), err
	}
	return "comment correctly posted", nil
}
