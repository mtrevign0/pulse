package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirestoreClient struct {
	client *firestore.Client
}

func NewFirestoreClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*FirestoreClient, error) {
	client, err := firestore.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	return &FirestoreClient{client: client}, nil
}
