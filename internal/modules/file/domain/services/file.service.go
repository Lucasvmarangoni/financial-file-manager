package services

import (
	"errors"
	"log"
	"os"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/repositories"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	zerolog "github.com/rs/zerolog/log"
	maliciouspdfdetector "github.com/saudshaddad/malicious-pdf-detector"

)

type FileService struct {
	Contract           *entities.Contract
	Extract            *entities.Extract
	Invoice            *entities.Invoice
	ContractRepository repositories.ContractRepository
}

func NewFileService() FileService{
	return FileService{}
}


func (fs *FileService) security(file []byte, dir string) error {

	fileName := fs.Contract.File.ID.String() + ".pdf"

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
	}

	err := os.WriteFile(dir+fileName, file, 0600)
	if err != nil {
		return err
	}

	PDFFile := maliciouspdfdetector.NewPDFFile(dir + fileName)

	err = PDFFile.ReadFile()
	if err != nil {
		return err
	}

	PDFFile.ParsePdfFile()

	if PDFFile.IsMalicious() {
		err = os.Remove(dir + fileName)
		if err != nil {
			zerolog.Fatal().Err(err).Str("Security", "file.service.go").Stack().Msg("Error to remove malicious file")
			return err
		}
		return errors.New("The file is probably malicious")
	}

	return nil
}

func (fs *FileService) optimization(dir string, outDir string, fileName string) ([]byte, error) {

	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		err := os.Mkdir(outDir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	err := api.OptimizeFile(dir, outDir, nil)
	if err != nil {
		log.Fatal(err)
	}

	optimizedFile, err := os.ReadFile(outDir)
	if err != nil {
		log.Fatal(err)
	}
	return optimizedFile, nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
