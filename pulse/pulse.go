package pulse

import (
	"context"

	"pulse/auth"
	"google.golang.org/api/option"
)

type Pulse struct {
	FirebaseAuth     *auth.FirebaseAuth
}

func NewPulse(ctx context.Context, projectID string, opts ...option.ClientOption) (*Pulse, error) {
	firebaseAuth, err := auth.NewFirebaseAuth(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &Pulse{
		FirebaseAuth:     firebaseAuth,
	}, nil
}
