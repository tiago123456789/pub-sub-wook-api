-- +goose Up
-- +goose StatementBegin
CREATE TABLE urls_producers(
    id SERIAL PRIMARY KEY,
    enabled BOOLEAN DEFAULT TRUE,
    key varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE urls_producers;
-- +goose StatementEnd
