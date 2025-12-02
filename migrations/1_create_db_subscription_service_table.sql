-- +goose Up
-- +goose StatementBegin
create table subscription (
   subscription_id text primary key,
   service_name text not null,
   price integer not null,
   user_id UUID default gen_random_uuid(),
   start_date timestamptz default now(),
   end_date timestamptz
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subscription;
-- +goose StatementEnd