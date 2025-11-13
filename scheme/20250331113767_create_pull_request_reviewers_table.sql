-- +goose Up
-- +goose StatementBegin


CREATE TABLE pull_request_reviewers (
                                        pull_request_id VARCHAR(128) NOT NULL,
                                        reviewer_id VARCHAR(64) NOT NULL,
                                        PRIMARY KEY (pull_request_id, reviewer_id)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE pull_request_reviewers CASCADE;
-- +goose StatementEnd