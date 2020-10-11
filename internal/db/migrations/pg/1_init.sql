-- +goose Up
CREATE TABLE IF NOT EXISTS production.users (
    id serial PRIMARY KEY,
    user_hash text,
    user_login character varying(255),
    user_email character varying(255),
    user_phone character varying(255),
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT user_login_unique UNIQUE (user_login),
    CONSTRAINT user_email_unique UNIQUE (user_email),
    CONSTRAINT user_phone_unique UNIQUE (user_phone),
);

CREATE TABLE IF NOT EXISTS production.profiles (
    id INTEGER PRIMARY KEY,
    user_first_name character varying(255),
    user_middle_name character varying(255),
    user_last_name character varying(255),
    user_position character varying(255),
    user_company character varying(255),
    user_private_key character varying(255),
    user_public_key character varying(255),
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS production.acts (
    id serial PRIMARY KEY,
    user_id INTEGER NOT NULL,
    finished boolean,
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

-- +goose Down
DROP TABLE production.users;
DROP TABLE production.profiles;
DROP TABLE production.acts;