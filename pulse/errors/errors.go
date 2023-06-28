package errors

import "fmt"

type ConnectionFailure struct {
	ServiceName string
}

func (e *ConnectionFailure) Error() string {
	return fmt.Sprintf("Failed to connect to service: %s", e.ServiceName)
}

func ConnectionFailureError(serviceName string) error {
	return &ConnectionFailure{ServiceName: serviceName}
}

type UnknownService struct {
	ServiceName string
}

func (e *UnknownService) Error() string {
	return fmt.Sprintf("Unknown service: %s", e.ServiceName)
}

func UnknownServiceError(serviceName string) error {
	return &UnknownService{ServiceName: serviceName}
}

type MethodFailure struct {
	ServiceName string
}

func (e *MethodFailure) Error() string {
	return fmt.Sprintf("Failed to invoke method for service: %s", e.ServiceName)
}

func MethodFailureError(serviceName string) error {
	return &MethodFailure{ServiceName: serviceName}
}

type PublishFailure struct {
	ServiceName string
}

func (e *PublishFailure) Error() string {
	return fmt.Sprintf("Failed to publish message for service: %s", e.ServiceName)
}

func PublishFailureError(serviceName string) error {
	return &PublishFailure{ServiceName: serviceName}
}
