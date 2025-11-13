-- +goose Up
-- +goose StatementBegin


CREATE TABLE pull_requests (
                               pull_request_id VARCHAR(128) PRIMARY KEY,
                               pull_request_name VARCHAR(255) NOT NULL,
                               author_id VARCHAR(64) NOT NULL, -- без FK, проверяешь в коде
                               status VARCHAR(16) NOT NULL CHECK (status IN ('OPEN', 'MERGED')),
                               created_at TIMESTAMP DEFAULT NOW(),
                               merged_at TIMESTAMP
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE pull_requests CASCADE;
-- +goose StatementEnd