# Pulse

Pulse is a powerful package for managing service connections and providing failover mechanisms using gRPC and PubSub in Go.

## Features

- Connect to gRPC services and manage client connections
- Automatic failover to PubSub messaging when gRPC connection fails
- Flexible and extensible design for easy integration

## Installation

To install the Pulse package, use the following `go get` command:

```bash
go get -u github.com/mtrevign0/pulse
```
## Usage
Let's use the `ServiceConnector`:

```go
package main

import (
	"context"
	"fmt"
	"pulse/service_connector"
)

func main() {
	sc := service_connector.NewServiceConnector()

	// Connect to a gRPC service
	err := sc.ConnectService(context.Background(), "example", "localhost:50051")
	if err != nil {
		fmt.Printf("Failed to connect to service: %v\n", err)
		return
	}

	// Call a gRPC method or failover to PubSub
	err = sc.CallServiceOrPublish(context.Background(), "example", func(conn *grpc.ClientConn) error {
		// Make a gRPC method call using the connection
		// ...
		return nil
	}, payload)

	if err != nil {
		fmt.Printf("Error during service call or failover: %v\n", err)
		return
	}
}
```
As you can see, this service connector handles synchronous message delivery to a gRPC service by dialing to it, and optionally, publish the message to a homonimous topic in Pub/Sub should the gRPC dial fail.

## Documentation
Full documentation is a Work in Progress.

## Contributing
Contributions are welcome, however as I'm not expecting any contributions soon, I haven't established the mechanism yet. In any case, if you'd wish to contribute, email me directly:
`mario@trevign.com`

## License
This project is licensed under the MIT License.
