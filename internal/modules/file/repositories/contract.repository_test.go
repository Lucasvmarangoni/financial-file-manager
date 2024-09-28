package repositories_test

import (
	"context"
	"testing"
	"time"

	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/domain/entities"
	"github.com/Lucasvmarangoni/financial-file-manager/internal/modules/file/repositories"
	pkg_entities "github.com/Lucasvmarangoni/financial-file-manager/pkg/entities"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func TestContractRepositoryDb_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTx := NewMockTx(ctrl)

	contract := &entities.Contract{
		File: entities.File{
			ID:        pkg_entities.NewID(),
			Type:      "type",
			CreatedAt: time.Now(),
			Customer:  "user",
		},
		Title:   "title",
		Parties: []string{"party1", "party2"},
		Object:  "object",
		Extract: []string{"extract1", "extract2"},
		Invoice: []string{"invoice1", "invoice2"},
	}

	ctx := context.Background()

	mockTx.EXPECT().Exec(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(pgx.CommandTag{}, nil)

	repo := repositories.NewContractRepository(mockTx)
	_, err := repo.Insert(contract, ctx, mockTx)

	assert.NoError(t, err)
}
