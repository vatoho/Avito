-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_banner_tags_feature_tag ON banner_tags (feature_id, tag_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_banner_tags_feature_tag;
-- +goose StatementEnd
