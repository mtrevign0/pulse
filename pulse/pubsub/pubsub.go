package pubsub

import (
	"context"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type PubSubClient struct {
	Client *pubsub.Client
}

type PubSubMessage struct {
	Data []byte
}

func NewPubSubClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*PubSubClient, error) {
	client, err := pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	return &PubSubClient{Client: client}, nil
}
