-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_previous_banners_banner_id ON previous_banners (banner_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_previous_banners_banner_id;
-- +goose StatementEnd
