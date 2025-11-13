-- +goose Up
-- +goose StatementBegin


CREATE TABLE teams (
                       team_name VARCHAR(255) PRIMARY KEY
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE teams CASCADE;
-- +goose StatementEnd