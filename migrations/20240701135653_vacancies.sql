-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS vacancies (
  vacancy varchar(255) NOT NULL,
  salary integer,
  email varchar(255),
  city varchar(255),
  user_id integer references users (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS vacancies;
-- +goose StatementEnd
