-- +goose Up
-- +goose StatementBegin
CREATE INDEX events_begin_time_idx ON events (begin_time);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX events_begin_time_idx;
-- +goose StatementEnd
