-- +goose Up
-- +goose StatementBegin
CREATE TABLE journey_tag (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE journey_tag;
-- +goose StatementEnd
