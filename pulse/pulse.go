package pulse

import (
	"context"

	"pulse/auth"
	"pulse/firestore"
	"pulse/pubsub"
	"pulse/service_connector"
	"google.golang.org/api/option"
)

type Pulse struct {
	FirebaseAuth     *auth.FirebaseAuth
	FirestoreClient  *firestore.FirestoreClient
	PubSubClient     *pubsub.PubSubClient
	ServiceConnector *service_connector.ServiceConnector
}

func NewPulse(ctx context.Context, projectID string, opts ...option.ClientOption) (*Pulse, error) {
	firebaseAuth, err := auth.NewFirebaseAuth(ctx, opts...)
	if err != nil {
		return nil, err
	}
	firestoreClient, err := firestore.NewFirestoreClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	pubsubClient, err := pubsub.NewPubSubClient(ctx, projectID, opts...)
	if err != nil {
		return nil, err
	}
	serviceConnector := service_connector.NewServiceConnector()

	return &Pulse{
		FirebaseAuth:     firebaseAuth,
		FirestoreClient:  firestoreClient,
		PubSubClient:     pubsubClient,
		ServiceConnector: serviceConnector,
	}, nil
}
