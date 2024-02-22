-- +goose Up
-- +goose StatementBegin
CREATE TABLE counters (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    ip_address VARCHAR(20),
    count MEDIUMINT NOT NULL DEFAULT 0
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE counters;
-- +goose StatementEnd
