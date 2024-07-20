package model

import "time"

type Mock struct {
	ID          uint              `json:"id" gorm:"primaryKey"`
	Method      string            `json:"method" gorm:"type:varchar(10);not null"`
	URL         string            `json:"url" gorm:"type:text;not null"`
	RequestBody string            `json:"request_body" gorm:"type:text"`
	StatusCode  int               `json:"status_code"`
	Headers     map[string]string `json:"headers"`
	Body        interface{}       `json:"body"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type MockResponse struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       interface{}       `json:"body"`
}
