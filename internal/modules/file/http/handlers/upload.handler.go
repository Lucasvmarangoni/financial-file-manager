package handlers

import (
	"context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/pb"
)

type Server struct {
	pb.UnimplementedContractRequestServer
	pb.UnimplementedExtractRequestServer
	pb.UnimplementedInvoiceRequestServer
}

func (s *Server) ContractRequest(ctx context.Context, in *pb.Contract) (*pb.Response, error) {
	return &pb.Response{Success: true, Message: "teste"}, nil
}

func (s *Server) ExtractRequest(ctx context.Context, in *pb.Extract) (*pb.Response, error) {	
	return &pb.Response{Success: true, Message: "teste"}, nil
}

func (s *Server) InvoiceRequest(ctx context.Context, in *pb.Invoice) (*pb.Response, error) {
	return &pb.Response{Success: true, Message: "teste"}, nil
}