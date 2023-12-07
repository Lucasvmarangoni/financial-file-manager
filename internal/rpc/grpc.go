package rpc

import (	
	"net"
	"os"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/http/handlers"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func Connect() {

	listener, err := net.Listen("tcp", "localhost:" + os.Getenv("GRPC_PORT"))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterContractRequestServer(grpcServer, &handlers.Server{})
	pb.RegisterExtractRequestServer(grpcServer, &handlers.Server{})
	pb.RegisterInvoiceRequestServer(grpcServer, &handlers.Server{})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal().Err(err).Stack().Time("time", time.Now()).Msg("Failed to start gRPC server")
	}
}
