-- +goose Up
-- +goose StatementBegin
CREATE TABLE "banners"
(
    id         SERIAL PRIMARY KEY NOT NULL,
    content    VARCHAR            NOT NULL,
    created_at TIMESTAMP          NOT NULL,
    updated_at TIMESTAMP          NOT NULL,
    is_active  BOOLEAN            NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE banners;
-- +goose StatementEnd
