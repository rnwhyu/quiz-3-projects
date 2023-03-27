-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE categories (
    id serial4 NOT NULL,
    "name" VARCHAR NOT NULL,
    created_date timestamp DEFAULT now() NOT NULL,
    updated_date timestamp DEFAULT now() NOT NULL
)

-- +migrate StatementEnd