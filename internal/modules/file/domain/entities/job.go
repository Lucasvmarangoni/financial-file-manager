package entities

import (
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Job struct {
	ID               entities.ID `json:"job_id" valid:"uuid" `
	OutputBucketPath string      `json:"output_bucket_path" valid:"notnull"`
	Status           string      `json:"status" valid:"notnull"`
	FileID           string      `valid:"-" `
	Error            string      `valid:"-"`
	CreatedAt        time.Time   `json:"created_at" valid:"-"`
	UpdatedAt        time.Time   `json:"updated_at" valid:"-"`
}

func NewJob(output string, status string) (*Job, error) {
	job := Job{
		OutputBucketPath: output,
		Status:           status,
	}
	job.prepare()

	err := job.Validate()
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (job *Job) prepare() {
	job.ID = entities.NewID()
	job.CreatedAt = time.Now()
	job.UpdatedAt = time.Now()
}

func (job *Job) Validate() error {
	_, err := govalidator.ValidateStruct(job)

	if err != nil {
		return err
	}
	return nil
}
