package tests

import (
	"context"
	"errors"
	"testing"
	"google.golang.org/grpc"
)

type MockServiceConnector struct{}

func (m *MockServiceConnector) ConnectService(ctx context.Context, name string, address string) error {
	if name == "" || address == "" {
		return errors.New("invalid arguments")
	}
	return nil
}

func (m *MockServiceConnector) GetService(name string) (*grpc.ClientConn, error) {
	if name == "" {
		return nil, errors.New("invalid argument")
	}
	return &grpc.ClientConn{}, nil
}

func TestConnectService(t *testing.T) {
	mock := &MockServiceConnector{}

	err := mock.ConnectService(context.Background(), "test", "localhost:50051")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = mock.GetService("test")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	_, err = mock.GetService("")
	if err == nil {
		t.Errorf("Expected an error for invalid argument")
	}
}
