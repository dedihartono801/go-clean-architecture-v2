package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/dedihartono801/go-clean-architecture-v2/database"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/repository"
	"github.com/dedihartono801/go-clean-architecture-v2/internal/app/usecase/grpc/transaction"
	grpcHandler "github.com/dedihartono801/go-clean-architecture-v2/internal/delivery/grpc"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/config"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/identifier"
	pb "github.com/dedihartono801/go-clean-architecture-v2/pkg/protobuf"
	"github.com/dedihartono801/go-clean-architecture-v2/pkg/validator"
	validatorv10 "github.com/go-playground/validator/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	envConfig := config.SetupEnvFile()
	mysql := database.InitMysql(envConfig)

	identifier := identifier.NewIdentifier()
	validator := validator.NewValidator(validatorv10.New())
	transactionRepository := repository.NewTransactionRepository(mysql)
	transactionService := transaction.NewGrpcTransactionService(transactionRepository, validator, identifier)
	transactionHandler := grpcHandler.Service{Service: transactionService}

	lis, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("GRPC Svc on", os.Getenv("GRPC_PORT"))

	opts := []grpc.ServerOption{}
	tls := false //in this example we do not use ssl

	// if using ssl server side
	if tls {
		certFile := "pkg/ssl/server.crt" //example ssl path
		kefFile := "pkg/ssl/server.pem"  //example ssl path

		creds, err := credentials.NewServerTLSFromFile(certFile, kefFile)

		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterTransactionServiceServer(grpcServer, &transactionHandler)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
