package handlers_test

import (
	"context"
	"testing"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/http/handlers"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/pb"
	"github.com/stretchr/testify/assert"
)

func TestServer_ContractRequest(t *testing.T) {
	s := handlers.Server{}
	resp, err := s.ContractRequest(context.Background(), &pb.Contract{})
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, "teste", resp.Message)
}

func TestServer_ExtractRequest(t *testing.T) {
	s := handlers.Server{}
	resp, err := s.ExtractRequest(context.Background(), &pb.Extract{})
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, "teste", resp.Message)
}

func TestServer_InvoiceRequest(t *testing.T) {
	s := handlers.Server{}
	resp, err := s.InvoiceRequest(context.Background(), &pb.Invoice{})
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	assert.Equal(t, "teste", resp.Message)
}