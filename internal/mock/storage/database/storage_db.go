package storage

import (
	"WbTest/internal/infrastructure/database/postgres/database"
	"WbTest/internal/mock/model"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"log"
)

const expireTime = 15

type MockStorageDB struct {
	db              database.Database
	cacheExpireTime int
	logger          *zap.SugaredLogger
}

func New(db database.Database, logger *zap.SugaredLogger) *MockStorageDB {
	return &MockStorageDB{
		db:     db,
		logger: logger,
	}
}

func (s *MockStorageDB) SaveMock(ctx context.Context, mock model.Mock) error {
	headers, err := json.Marshal(mock.Headers)
	if err != nil {
		return fmt.Errorf("failed to marshal headers: %w", err)
	}

	body, err := json.Marshal(mock.Body)
	if err != nil {
		return fmt.Errorf("failed to marshal body: %w", err)
	}

	query := `
		INSERT INTO mocks (method, url, request_body, status_code, headers, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id;
	`

	_, err = s.db.Exec(ctx, query,
		mock.Method,
		mock.URL,
		mock.RequestBody,
		mock.StatusCode,
		headers,
		body,
		mock.CreatedAt,
		mock.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save mock: %w", err)
	}

	return nil
}

func (s *MockStorageDB) GetMocks(ctx context.Context) ([]model.Mock, error) {
	query := `SELECT id, method, url, request_body, status_code, headers, body, created_at, updated_at FROM mocks`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get mocks: %w", err)
	}
	defer rows.Close()

	var mocks []model.Mock
	for rows.Next() {
		var mock model.Mock
		var headers, body []byte
		if err := rows.Scan(&mock.ID, &mock.Method, &mock.URL, &mock.RequestBody, &mock.StatusCode, &headers, &body, &mock.CreatedAt, &mock.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan mock: %w", err)
		}
		if err := json.Unmarshal(headers, &mock.Headers); err != nil {
			return nil, fmt.Errorf("failed to unmarshal headers: %w", err)
		}
		if err := json.Unmarshal(body, &mock.Body); err != nil {
			return nil, fmt.Errorf("failed to unmarshal body: %w", err)
		}
		mocks = append(mocks, mock)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return mocks, nil
}

// GetMockByMethodAndURL возвращает мок по методу и URL.

func (s *MockStorageDB) GetMockByMethodAndURL(ctx context.Context, method string, url string) (*model.Mock, error) {
	query := `SELECT id, method, url, request_body, status_code, headers, body, created_at, updated_at FROM mocks WHERE method = $1 AND url = $2`
	log.Printf("Executing query: %s with method: %s, url: %s", query, method, url)

	row := s.db.QueryRow(ctx, query, method, url)

	var mock model.Mock
	var headers, body []byte
	if err := row.Scan(&mock.ID, &mock.Method, &mock.URL, &mock.RequestBody, &mock.StatusCode, &headers, &body, &mock.CreatedAt, &mock.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no mock found for method %s and url %s", method, url)
		}
		return nil, fmt.Errorf("failed to scan mock: %w", err)
	}
	if err := json.Unmarshal(headers, &mock.Headers); err != nil {
		return nil, fmt.Errorf("failed to unmarshal headers: %w", err)
	}
	if err := json.Unmarshal(body, &mock.Body); err != nil {
		return nil, fmt.Errorf("failed to unmarshal body: %w", err)
	}

	return &mock, nil
}
