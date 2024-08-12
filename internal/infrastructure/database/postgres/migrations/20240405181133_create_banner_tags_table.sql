-- +goose Up
-- +goose StatementBegin
CREATE TABLE "banner_tags"
(
    feature_id INT NOT NULL,
    tag_id     INT NOT NULL,
    banner_id  INT NOT NULL REFERENCES banners (id) ON DELETE CASCADE,

    PRIMARY KEY (feature_id, tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE banner_tags;
-- +goose StatementEnd
