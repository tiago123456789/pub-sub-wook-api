package config

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
)

type Message struct {
	Queue   string
	Message string
}

type IQueue interface {
	SendMessage(params Message) error
}

type Queue struct {
	client *sqs.Client
}

func NewQueue() *Queue {
	httpClient := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        10,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	awsProfile := os.Getenv("AWS_PROFILE")
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithSharedConfigProfile(awsProfile),
		config.WithHTTPClient(httpClient),
	)
	if err != nil {
		log.Fatal("unable to load SDK config, " + err.Error())
	}

	client := sqs.NewFromConfig(cfg)

	return &Queue{
		client: client,
	}
}

func (queue *Queue) SendMessage(params Message) error {
	_, err := queue.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		MessageBody: &params.Message,
		QueueUrl:    aws.String(params.Queue),
	})

	return err
}

// func SendMessage(client *sqs.Client, queueUrl string, messageBody string) error {
// 	_, err := client.SendMessage(context.TODO(), &sqs.SendMessageInput{
// 		MessageBody: &messageBody,
// 		QueueUrl:    aws.String(queueUrl),
// 	})

// 	return err
// }
