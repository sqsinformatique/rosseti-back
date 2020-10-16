-- +goose Up
CREATE TABLE IF NOT EXISTS production.users (
    id serial PRIMARY KEY,
    user_hash text,
    user_email character varying(255),
    user_phone character varying(255),
    user_role character varying(255),
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT user_email_unique UNIQUE (user_email),
    CONSTRAINT user_phone_unique UNIQUE (user_phone)
);

CREATE TABLE IF NOT EXISTS production.profiles (
    id INTEGER PRIMARY KEY,
    user_first_name character varying(255),
    user_middle_name character varying(255),
    user_last_name character varying(255),
    user_position character varying(255),
    user_company character varying(255),
    user_private_key character varying(2048),
    user_public_key character varying(2048),
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS production.acts (
    id serial PRIMARY KEY,
    staff_id INTEGER NOT NULL,
    superviser_id INTEGER NOT NULL,
    object_id INTEGER NOT NULL,
    review_id INTEGER NOT NULL,
    finished boolean,
    end_at timestamp with time zone,
    staff_sign character varying(2048), 
    superviser_sign character varying(2048),
    approved boolean,
    reverted boolean,
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);



CREATE TABLE IF NOT EXISTS production.acts_details (
    act_id INTEGER NOT NULL,
    element_id INTEGER NOT NULL,
    defects jsonb,
    category INTEGER,
    repaired_at timestamp with time zone,
    images jsonb,
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS production.orders (
    id serial PRIMARY KEY,
    object_id INTEGER NOT NULL,
    tech_tasks jsonb,
    superviser_id INTEGER NOT NULL,
    staff_id INTEGER NOT NULL,
    start_at timestamp with time zone,
    end_at timestamp with time zone,
    superviser_sign character varying(2048),
    staff_sign character varying(2048),
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS production.tech_tasks (
    id serial PRIMARY KEY,
    description character varying(2048),
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS production.objects (
    id serial PRIMARY KEY,
    object_address character varying(2048),
    object_name character varying(2048),
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS production.objects_details (
    object_id INTEGER NOT NULL,
    element_id INTEGER NOT NULL,
    element_name character varying(2048),
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

CREATE TABLE IF NOT EXISTS production.sessions (
    id character varying(255) PRIMARY KEY,
    user_id INTEGER NOT NULL,
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

-- +goose Down
DROP TABLE production.users;
DROP TABLE production.profiles;
DROP TABLE production.acts;
DROP TABLE production.orders;
DROP TABLE production.objects;