-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS mocks (
    id SERIAL PRIMARY KEY,
    method VARCHAR(10) NOT NULL,
    url TEXT NOT NULL,
    request_body TEXT,
    status_code INTEGER NOT NULL,
    headers TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
                             );


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS mocks;
-- +goose StatementEnd
