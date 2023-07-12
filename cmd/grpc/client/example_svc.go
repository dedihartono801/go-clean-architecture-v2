package client

import (
	"fmt"
	"log"

	"github.com/dedihartono801/go-clean-architecture-v2/pkg/config"
	pb "github.com/dedihartono801/go-clean-architecture-v2/pkg/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// example client
func InitExampleServiceClient(c *config.Config) pb.ExampleServiceClient {
	opts := []grpc.DialOption{}
	tls := false

	// example connect with no ssl
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if tls {
		certFile := "pkg/ssl/ca.crt" //example ca.crt path

		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificates: %v\n", err)
		}

		// example connect with ssl
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	cc, err := grpc.Dial(c.ExamplSvcUrl, opts...)

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewExampleServiceClient(cc)
}
