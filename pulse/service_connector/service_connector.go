package service_connector

import (
	"context"
	"sync"

	"pulse/errors"
	"google.golang.org/grpc"
)

type ServiceConnector struct {
	connections map[string]*grpc.ClientConn
	mu          sync.Mutex
}

func NewServiceConnector() *ServiceConnector {
	return &ServiceConnector{connections: make(map[string]*grpc.ClientConn)}
}

func (s *ServiceConnector) ConnectService(ctx context.Context, name string, address string) error {
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	s.mu.Lock()
	s.connections[name] = conn
	s.mu.Unlock()
	return nil
}

func (s *ServiceConnector) GetService(name string) (*grpc.ClientConn, error) {
	s.mu.Lock()
	conn, ok := s.connections[name]
	s.mu.Unlock()
	if !ok {
		return nil, errors.ErrServiceNotFound
	}
	return conn, nil
}
