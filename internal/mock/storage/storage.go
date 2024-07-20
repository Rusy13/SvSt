package storage

import (
	"WbTest/internal/mock/model"
	"context"
)

type MockStorage interface {
	SaveMock(ctx context.Context, mock model.Mock) error
	GetMocks(ctx context.Context) ([]model.Mock, error)
	GetMockByMethodAndURL(ctx context.Context, method string, url string) (*model.Mock, error)
}
