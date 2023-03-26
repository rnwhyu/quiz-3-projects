-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE books(
    id integer NOT NULL,
    title character varying NOT NULL,
    description character varying NOT NULL,
    image_url character varying NOT NULL,
    release_year integer NOT NULL,
    price character varying NOT NULL,
    total_page integer NOT NULL,
    thickness character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    category_id integer NOT NULL
);
CREATE TABLE categories (
    id integer NOT NULL,
    name character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);

-- +migrate StatementEnd