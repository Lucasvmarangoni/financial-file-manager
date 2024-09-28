package services

// import (
// 	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
// 	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/repositories"
// 	"github.com/Lucasvmarangoni/financial-file-manager/config"
// )

// type JobService struct {
// 	Job           *entities.Job
// 	JobRepository repositories.JobRepository
// 	FileService  FileService
// }

// func (j *JobService) StartJob() error {

// 	storage_local_path := config.GetEnvString("storage_local_path")

// 	err := j.changeJobStatus("SECURITY VERIFY")
// 	if err != nil {
// 		return j.failJob(err)
// 	}

// 	err = j.FileService.security(, storage_local_path)
// 	if err != nil {
// 		return j.failJob(err)
// 	}

// 	err = j.changeJobStatus("FRAGMENTING")
// 	if err != nil {
// 		return j.failJob(err)
// 	}

// 	err = j.FileService.Fragment()
// 	if err != nil {
// 		return j.failJob(err)
// 	}

// 	err = j.changeJobStatus("ENCODING")
// 	if err != nil {
// 		return j.failJob(err)
// 	}

// 	err = j.FileService.Encode()
// 	if err != nil {
// 		return j.failJob(err)
// 	}

// 	err = j.performUpload()
// 	if err != nil {
// 		return j.failJob(err)
// 	}

// 	err = j.changeJobStatus("FINISHING")
// 	if err != nil {
// 		return j.failJob(err)
// 	}

// 	err = j.FileService.Finish()
// 	if err != nil {
// 		return j.failJob(err)
// 	}

// 	err = j.changeJobStatus("COMPLETED")
// 	if err != nil {
// 		return j.failJob(err)
// 	}

// 	return nil
// }
