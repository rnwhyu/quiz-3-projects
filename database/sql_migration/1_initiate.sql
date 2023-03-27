-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE books(
    book_id serial4 NOT NULL,
    book_title VARCHAR NOT NULL,
    book_desc VARCHAR NOT NULL,
    img_url VARCHAR NOT NULL,
    book_year integer NOT NULL,
    book_price VARCHAR NOT NULL,
    book_page integer NOT NULL,
    book_thick VARCHAR NOT NULL,
    created_date timestamp DEFAULT now() NOT NULL,
    updated_date timestamp DEFAULT now() NOT NULL,
    category_id integer NOT NULL
)
-- +migrate StatementEnd