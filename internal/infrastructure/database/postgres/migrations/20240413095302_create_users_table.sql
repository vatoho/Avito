-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users"
(
    tag_id INT PRIMARY KEY    NOT NULL,
    token  VARCHAR(64) UNIQUE NOT NULL,
    role   VARCHAR(32)        NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users"
-- +goose StatementEnd
