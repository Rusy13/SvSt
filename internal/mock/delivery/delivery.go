package delivery

import (
	"WbTest/internal/mock/service"
	"go.uber.org/zap"
)

type MockDelivery struct {
	service service.MockService
	logger  *zap.SugaredLogger
}

func New(service service.MockService, logger *zap.SugaredLogger) *MockDelivery {
	return &MockDelivery{
		service: service,
		logger:  logger,
	}
}
