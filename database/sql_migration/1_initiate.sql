-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE books(
    id serial4 NOT NULL,
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    image_url VARCHAR NOT NULL,
    release_year INT NOT NULL,
    price VARCHAR NOT NULL,
    total_page INT NOT NULL,
    thickness VARCHAR NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL,
    category_id integer NOT NULL
);
CREATE TABLE categories (
    id serial4 NOT NULL,
    "name" VARCHAR NOT NULL,
    created_at timestamp DEFAULT now() NOT NULL,
    updated_at timestamp DEFAULT now() NOT NULL
);

-- +migrate StatementEnd