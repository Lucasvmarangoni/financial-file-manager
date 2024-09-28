package services

import (
	// "context"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/pb"
)

type JobManager struct {
	ContractData   *pb.ContractData
	ExtractData    *pb.ExtractData
	InvoiceData    *pb.InvoiceData
	File           []byte		
}

type JobNotificationError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func NewJobManager(
	contractData *pb.ContractData,
	extractData *pb.ExtractData,
	invoiceData *pb.InvoiceData,
	file []byte,
) *JobManager {
	return &JobManager{
		ContractData:   contractData,
		ExtractData:    extractData,
		InvoiceData:    invoiceData,
		File:           file,		
	}
}


// func (j *JobManager) Start(ctx context.Context) error {

// }
