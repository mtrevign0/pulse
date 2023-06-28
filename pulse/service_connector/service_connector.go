package service_connector

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"time"
	"pulse/errors"
	"pulse/pubsub"
	cloudpubsub "cloud.google.com/go/pubsub"
)

type ServiceConnector struct {
	services map[string]*grpc.ClientConn
	PubSubClient *pubsub.PubSubClient
}

func NewServiceConnector() *ServiceConnector {
	return &ServiceConnector{
		services: make(map[string]*grpc.ClientConn),
	}
}

func (sc *ServiceConnector) ConnectService(ctx context.Context, name string, address string) error {
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		return errors.ConnectionFailureError(name)
	}
	sc.services[name] = conn
	return nil
}

func (sc *ServiceConnector) GetService(name string) (*grpc.ClientConn, error) {
	conn, ok := sc.services[name]
	if !ok {
		return nil, errors.UnknownServiceError(name)
	}
	return conn, nil
}

func (sc *ServiceConnector) CallServiceOrPublish(ctx context.Context, name string, call func(conn *grpc.ClientConn) error, payload interface{}) error {
	conn, err := sc.GetService(name)
	if err != nil {
		return err
	}

	err = call(conn)
	if err != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return errors.MethodFailureError(name)
		}

		t := sc.PubSubClient.Client.Topic(name)
		res := t.Publish(ctx, &cloudpubsub.Message{Data: data})
		if _, err = res.Get(ctx); err != nil {
			return errors.PublishFailureError(name)
		}
	}

	return nil
}