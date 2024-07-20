package service

import (
	"context"
	"fmt"
	"log"

	"WbTest/internal/mock/model"
	"WbTest/internal/mock/storage"
)

// WeatherService определяет методы для работы с данными о погоде.
type MockService interface {
	CreateMock(ctx context.Context, mock *model.Mock) error
	GetAllMocks(ctx context.Context) ([]model.Mock, error)
	GetMock(ctx context.Context, method, url string) (*model.Mock, error)
}

// MockServiceImpl реализует интерфейс MockService.
type MockServiceImpl struct {
	storage storage.MockStorage
}

// NewMockService создает новый экземпляр MockServiceImpl.
func NewMockService(storage storage.MockStorage) *MockServiceImpl {
	return &MockServiceImpl{
		storage: storage,
	}
}

// CreateMock создает новый мок.
func (s *MockServiceImpl) CreateMock(ctx context.Context, mock *model.Mock) error {
	err := s.storage.SaveMock(ctx, *mock)
	if err != nil {
		// Логируем ошибку с дополнительной информацией
		log.Printf("Failed to save mock: %v", err)
		return fmt.Errorf("failed to save mock: %w", err)
	}
	return nil
}

// GetAllMocks возвращает все моки.
func (s *MockServiceImpl) GetAllMocks(ctx context.Context) ([]model.Mock, error) {
	mocks, err := s.storage.GetMocks(ctx)
	if err != nil {
		log.Printf("Failed to get mocks: %v", err)
		return nil, fmt.Errorf("failed to get mocks: %w", err)
	}
	return mocks, nil
}

func (s *MockServiceImpl) GetMock(ctx context.Context, method, url string) (*model.Mock, error) {
	mock, err := s.storage.GetMockByMethodAndURL(ctx, method, url)
	if err != nil {
		log.Printf("Failed to get mock for method: %s, url: %s, error: %v", method, url, err)
		return nil, fmt.Errorf("failed to get mock: %w", err)
	}
	return mock, nil
}
