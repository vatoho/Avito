-- +goose Up
-- +goose StatementBegin
CREATE TABLE "previous_banners"
(
    id         SERIAL PRIMARY KEY NOT NULL,
    banner_id  INT                NOT NULL REFERENCES banners (id) ON DELETE CASCADE,
    content    VARCHAR            NOT NULL,
    updated_at TIMESTAMP          NOT NULL,
    is_active  BOOLEAN            NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE previous_banners;
-- +goose StatementEnd
