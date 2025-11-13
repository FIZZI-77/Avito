-- +goose Up
-- +goose StatementBegin



CREATE TABLE users (
                       user_id VARCHAR(64) PRIMARY KEY,
                       username VARCHAR(255) NOT NULL,
                       team_name VARCHAR(255),
                       is_active BOOLEAN NOT NULL DEFAULT TRUE
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE users CASCADE;
-- +goose StatementEnd