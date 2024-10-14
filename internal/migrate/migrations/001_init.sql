-- see for more examples: https://github.com/rubenv/sql-migrate?tab=readme-ov-file#writing-migrations

-- +migrate Up
CREATE TABLE IF NOT EXISTS key_value (
       key   VARCHAR(200) PRIMARY KEY,
       value TEXT
);

-- +migrate Down
