package auth

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

type FirebaseAuth struct {
	client *auth.Client
}

func NewFirebaseAuth(ctx context.Context, opts ...option.ClientOption) (*FirebaseAuth, error) {
	app, err := firebase.NewApp(ctx, nil, opts...)
	if err != nil {
		return nil, err
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	return &FirebaseAuth{client: client}, nil
}
