-- +goose Up
CREATE TABLE IF NOT EXISTS production.users (
    id serial PRIMARY KEY,
    user_hash character varying(255) DEFAULT '',
    user_email character varying(255) DEFAULT '',
    user_phone character varying(255) DEFAULT '',
    user_role character varying(255) DEFAULT '',
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT user_email_unique UNIQUE (user_email),
    CONSTRAINT user_phone_unique UNIQUE (user_phone)
);

INSERT INTO production.users (user_hash, user_email, user_phone, user_role) 
VALUES ('15e2b0d3c33891ebb0f1ef609ec419420c20e320ce94c65fbc8c3312448eb225', 'test1@rosseti.ru', '+79169999999', 'MASTER');
INSERT INTO production.users (user_hash, user_email, user_phone, user_role) 
VALUES ('15e2b0d3c33891ebb0f1ef609ec419420c20e320ce94c65fbc8c3312448eb225', 'test2@rosseti.ru', '+79169999998', 'ELECTRICIAN');
INSERT INTO production.users (user_hash, user_email, user_phone, user_role) 
VALUES ('15e2b0d3c33891ebb0f1ef609ec419420c20e320ce94c65fbc8c3312448eb225', 'test3@rosseti.ru', '+79169999997', 'ENGINEER');
INSERT INTO production.users (user_hash, user_email, user_phone, user_role) 
VALUES ('15e2b0d3c33891ebb0f1ef609ec419420c20e320ce94c65fbc8c3312448eb225', 'test4@rosseti.ru', '+79169999996', 'ADMIN');

CREATE TABLE IF NOT EXISTS production.profiles (
    id INTEGER PRIMARY KEY,
    user_first_name character varying(255) DEFAULT '',
    user_middle_name character varying(255) DEFAULT '',
    user_last_name character varying(255) DEFAULT '',
    user_position character varying(255) DEFAULT '',
    user_company character varying(255) DEFAULT '',
    user_private_key character varying(2048) DEFAULT '',
    user_public_key character varying(2048) DEFAULT '',
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

INSERT INTO production.profiles (id, user_first_name, user_middle_name, user_last_name, user_position, user_company, user_private_key, user_public_key) 
VALUES (1, 'Василий', 'Васильевич', 'Васильев', 'Мастер', 'ПАО МОЭСК', 
'MIIEpAIBAAKCAQEAsGSYKiaHzXTqMcIIdGE5iDU/3dlEVNMeUogwFSu83DN9JVNsmlIyhRVmEnm55sB6bs4ES0Lph28EeSmwf71sLWmmtWCxp90Cq3gBzDTzQJqyPNfL3eeFeeNDgNtUs3osB4gMdWs+CvbphRbOsAq3sbNQJ00aGZ6p6zkX/IZxrBIjxY+kxWIjQLy0p/Yn4rybM4+VFKMEUCiXM7Deu16N5hz1FF8HtSxsQBc4gszGWvIaVhZ1iK3y7DRqQBpjHCaA6cKC3Nh6ZpTxyIUphKrusBowYnXP5R/e3FQDzh5EhIwzIvEijJx9ywoeF9EC36Tnz62CRUIPbfla09oVlwGA1wIDAQABAoIBAQCDmons6NJpd9FDToEAU4mZFiGQY4mXv+vfp7w4D2nY4JF+R7+/Y5RNtqlxH2CTyQePpCWQAVw6r5mmzHPi2nDbcPfwWzQxCbP0OpUcxmS2zrQssNRpu1LanbS/buTDA2PWOqsQ7/JaO93+bgXHUje7XQ1wRRY0Byy/UtmSjrxApAqMiFGM4koIu9aOl6CvSb2bxO4MwitV9iS1lMFl4SlaYf9zPtIOJYUQWziRNz1ZK0CKCVnashspM+eNgVw+Z6P5OTOfOPDWBL0ORMvRXLQtoYND8tO9oCS+WV7MaeOk1OP8mP8okrb1o9x9ZGhWIQx7Ue7iaYvzgrUuj0gRtLABAoGBAN+PrpHnaY5DIgGqwBBTzNCaXX1W7n0wALU8qin5jo0cScSttMcM9OLM86BRMSCqYfOPDM0Lak3ZfgybxPCivhwwapsJ7ScP43/BzIFhzIZ2Xtg4P4KrF0OJ/ggfOUV6wgv0OKTzgTQ855lZy3MSKmVb7folXf9obYk+93PzZChvAoGBAMn80SrhuHZF81THBZfbcpJxpKNgTQerX4ia6e/CTT4hjlza8kRGawnWfQOFdmCT3mzdb6tQ+KPsxQWCRtH5DWg84GH/6+lgIUMEB+8OjUoDq6y+XmMWlCePVH1+7swHe7LRlaonwagA7SVn5t8+r5sZc+sIpH1wsc35DOZ+GlIZAoGAd3tBH3WAcqnqeN2bPJ6s7igyIxTc7UdEeZhskXZw+3XM7zKvVVrVXomPA3WhPgYRx6wCeWvKasT8mxx9SuaPmF0//JB3kNLrEZKwC84LEyocUo7tUpbCHjSX8htN7pZHM0BZLb9+pD6QwOK+20cwJW/WZkSmUiSrthhTBENmmj0CgYEAlhIUhju2hYlrRO2ppi4RbeSpYglGshANpr0SWmSOZz8fOrYhkcCP/nsx3s/mJ9M1SsUrFqnOUlyz9WfZnl/gKjYwsB8o8/fMPrJcAq1ZJEid4HaAQjagVNQU/ji0yzo0GaPGAuoO4/fsOgJ8chls91tt2I5PSDPWpyYHA6llfOECgYAhU0FLx2eHwQpFkQyGjgAvnRIzDwB6lKtFwO2BULwu8RBwYgQg1cAr8BCESJkpAHYC4FXEN1PT4qXb3fohGWOoIDWB6LliIUiGEFW5/1rFdLVUnXjA0NCqPQngBE/pW8/tdRnfqkjL07BRbz0RyLa5Me/AZfURkl21flBhvvs7bA==',
'MIIBCgKCAQEAsGSYKiaHzXTqMcIIdGE5iDU/3dlEVNMeUogwFSu83DN9JVNsmlIyhRVmEnm55sB6bs4ES0Lph28EeSmwf71sLWmmtWCxp90Cq3gBzDTzQJqyPNfL3eeFeeNDgNtUs3osB4gMdWs+CvbphRbOsAq3sbNQJ00aGZ6p6zkX/IZxrBIjxY+kxWIjQLy0p/Yn4rybM4+VFKMEUCiXM7Deu16N5hz1FF8HtSxsQBc4gszGWvIaVhZ1iK3y7DRqQBpjHCaA6cKC3Nh6ZpTxyIUphKrusBowYnXP5R/e3FQDzh5EhIwzIvEijJx9ywoeF9EC36Tnz62CRUIPbfla09oVlwGA1wIDAQAB'
);
INSERT INTO production.profiles (id, user_first_name, user_middle_name, user_last_name, user_position, user_company, user_private_key, user_public_key) 
VALUES (2, 'Пётр', 'Петрович', 'Петров', 'Электромонтёр', 'ПАО МОЭСК', 
'MIIEpAIBAAKCAQEAsGSYKiaHzXTqMcIIdGE5iDU/3dlEVNMeUogwFSu83DN9JVNsmlIyhRVmEnm55sB6bs4ES0Lph28EeSmwf71sLWmmtWCxp90Cq3gBzDTzQJqyPNfL3eeFeeNDgNtUs3osB4gMdWs+CvbphRbOsAq3sbNQJ00aGZ6p6zkX/IZxrBIjxY+kxWIjQLy0p/Yn4rybM4+VFKMEUCiXM7Deu16N5hz1FF8HtSxsQBc4gszGWvIaVhZ1iK3y7DRqQBpjHCaA6cKC3Nh6ZpTxyIUphKrusBowYnXP5R/e3FQDzh5EhIwzIvEijJx9ywoeF9EC36Tnz62CRUIPbfla09oVlwGA1wIDAQABAoIBAQCDmons6NJpd9FDToEAU4mZFiGQY4mXv+vfp7w4D2nY4JF+R7+/Y5RNtqlxH2CTyQePpCWQAVw6r5mmzHPi2nDbcPfwWzQxCbP0OpUcxmS2zrQssNRpu1LanbS/buTDA2PWOqsQ7/JaO93+bgXHUje7XQ1wRRY0Byy/UtmSjrxApAqMiFGM4koIu9aOl6CvSb2bxO4MwitV9iS1lMFl4SlaYf9zPtIOJYUQWziRNz1ZK0CKCVnashspM+eNgVw+Z6P5OTOfOPDWBL0ORMvRXLQtoYND8tO9oCS+WV7MaeOk1OP8mP8okrb1o9x9ZGhWIQx7Ue7iaYvzgrUuj0gRtLABAoGBAN+PrpHnaY5DIgGqwBBTzNCaXX1W7n0wALU8qin5jo0cScSttMcM9OLM86BRMSCqYfOPDM0Lak3ZfgybxPCivhwwapsJ7ScP43/BzIFhzIZ2Xtg4P4KrF0OJ/ggfOUV6wgv0OKTzgTQ855lZy3MSKmVb7folXf9obYk+93PzZChvAoGBAMn80SrhuHZF81THBZfbcpJxpKNgTQerX4ia6e/CTT4hjlza8kRGawnWfQOFdmCT3mzdb6tQ+KPsxQWCRtH5DWg84GH/6+lgIUMEB+8OjUoDq6y+XmMWlCePVH1+7swHe7LRlaonwagA7SVn5t8+r5sZc+sIpH1wsc35DOZ+GlIZAoGAd3tBH3WAcqnqeN2bPJ6s7igyIxTc7UdEeZhskXZw+3XM7zKvVVrVXomPA3WhPgYRx6wCeWvKasT8mxx9SuaPmF0//JB3kNLrEZKwC84LEyocUo7tUpbCHjSX8htN7pZHM0BZLb9+pD6QwOK+20cwJW/WZkSmUiSrthhTBENmmj0CgYEAlhIUhju2hYlrRO2ppi4RbeSpYglGshANpr0SWmSOZz8fOrYhkcCP/nsx3s/mJ9M1SsUrFqnOUlyz9WfZnl/gKjYwsB8o8/fMPrJcAq1ZJEid4HaAQjagVNQU/ji0yzo0GaPGAuoO4/fsOgJ8chls91tt2I5PSDPWpyYHA6llfOECgYAhU0FLx2eHwQpFkQyGjgAvnRIzDwB6lKtFwO2BULwu8RBwYgQg1cAr8BCESJkpAHYC4FXEN1PT4qXb3fohGWOoIDWB6LliIUiGEFW5/1rFdLVUnXjA0NCqPQngBE/pW8/tdRnfqkjL07BRbz0RyLa5Me/AZfURkl21flBhvvs7bA==',
'MIIBCgKCAQEAsGSYKiaHzXTqMcIIdGE5iDU/3dlEVNMeUogwFSu83DN9JVNsmlIyhRVmEnm55sB6bs4ES0Lph28EeSmwf71sLWmmtWCxp90Cq3gBzDTzQJqyPNfL3eeFeeNDgNtUs3osB4gMdWs+CvbphRbOsAq3sbNQJ00aGZ6p6zkX/IZxrBIjxY+kxWIjQLy0p/Yn4rybM4+VFKMEUCiXM7Deu16N5hz1FF8HtSxsQBc4gszGWvIaVhZ1iK3y7DRqQBpjHCaA6cKC3Nh6ZpTxyIUphKrusBowYnXP5R/e3FQDzh5EhIwzIvEijJx9ywoeF9EC36Tnz62CRUIPbfla09oVlwGA1wIDAQAB'
);
INSERT INTO production.profiles (id, user_first_name, user_middle_name, user_last_name, user_position, user_company, user_private_key, user_public_key) 
VALUES (3, 'Геннадий', 'Генадьевич', 'Геннадьев', 'Главный инженер', 'ПАО МОЭСК', 
'MIIEpAIBAAKCAQEAsGSYKiaHzXTqMcIIdGE5iDU/3dlEVNMeUogwFSu83DN9JVNsmlIyhRVmEnm55sB6bs4ES0Lph28EeSmwf71sLWmmtWCxp90Cq3gBzDTzQJqyPNfL3eeFeeNDgNtUs3osB4gMdWs+CvbphRbOsAq3sbNQJ00aGZ6p6zkX/IZxrBIjxY+kxWIjQLy0p/Yn4rybM4+VFKMEUCiXM7Deu16N5hz1FF8HtSxsQBc4gszGWvIaVhZ1iK3y7DRqQBpjHCaA6cKC3Nh6ZpTxyIUphKrusBowYnXP5R/e3FQDzh5EhIwzIvEijJx9ywoeF9EC36Tnz62CRUIPbfla09oVlwGA1wIDAQABAoIBAQCDmons6NJpd9FDToEAU4mZFiGQY4mXv+vfp7w4D2nY4JF+R7+/Y5RNtqlxH2CTyQePpCWQAVw6r5mmzHPi2nDbcPfwWzQxCbP0OpUcxmS2zrQssNRpu1LanbS/buTDA2PWOqsQ7/JaO93+bgXHUje7XQ1wRRY0Byy/UtmSjrxApAqMiFGM4koIu9aOl6CvSb2bxO4MwitV9iS1lMFl4SlaYf9zPtIOJYUQWziRNz1ZK0CKCVnashspM+eNgVw+Z6P5OTOfOPDWBL0ORMvRXLQtoYND8tO9oCS+WV7MaeOk1OP8mP8okrb1o9x9ZGhWIQx7Ue7iaYvzgrUuj0gRtLABAoGBAN+PrpHnaY5DIgGqwBBTzNCaXX1W7n0wALU8qin5jo0cScSttMcM9OLM86BRMSCqYfOPDM0Lak3ZfgybxPCivhwwapsJ7ScP43/BzIFhzIZ2Xtg4P4KrF0OJ/ggfOUV6wgv0OKTzgTQ855lZy3MSKmVb7folXf9obYk+93PzZChvAoGBAMn80SrhuHZF81THBZfbcpJxpKNgTQerX4ia6e/CTT4hjlza8kRGawnWfQOFdmCT3mzdb6tQ+KPsxQWCRtH5DWg84GH/6+lgIUMEB+8OjUoDq6y+XmMWlCePVH1+7swHe7LRlaonwagA7SVn5t8+r5sZc+sIpH1wsc35DOZ+GlIZAoGAd3tBH3WAcqnqeN2bPJ6s7igyIxTc7UdEeZhskXZw+3XM7zKvVVrVXomPA3WhPgYRx6wCeWvKasT8mxx9SuaPmF0//JB3kNLrEZKwC84LEyocUo7tUpbCHjSX8htN7pZHM0BZLb9+pD6QwOK+20cwJW/WZkSmUiSrthhTBENmmj0CgYEAlhIUhju2hYlrRO2ppi4RbeSpYglGshANpr0SWmSOZz8fOrYhkcCP/nsx3s/mJ9M1SsUrFqnOUlyz9WfZnl/gKjYwsB8o8/fMPrJcAq1ZJEid4HaAQjagVNQU/ji0yzo0GaPGAuoO4/fsOgJ8chls91tt2I5PSDPWpyYHA6llfOECgYAhU0FLx2eHwQpFkQyGjgAvnRIzDwB6lKtFwO2BULwu8RBwYgQg1cAr8BCESJkpAHYC4FXEN1PT4qXb3fohGWOoIDWB6LliIUiGEFW5/1rFdLVUnXjA0NCqPQngBE/pW8/tdRnfqkjL07BRbz0RyLa5Me/AZfURkl21flBhvvs7bA==',
'MIIBCgKCAQEAsGSYKiaHzXTqMcIIdGE5iDU/3dlEVNMeUogwFSu83DN9JVNsmlIyhRVmEnm55sB6bs4ES0Lph28EeSmwf71sLWmmtWCxp90Cq3gBzDTzQJqyPNfL3eeFeeNDgNtUs3osB4gMdWs+CvbphRbOsAq3sbNQJ00aGZ6p6zkX/IZxrBIjxY+kxWIjQLy0p/Yn4rybM4+VFKMEUCiXM7Deu16N5hz1FF8HtSxsQBc4gszGWvIaVhZ1iK3y7DRqQBpjHCaA6cKC3Nh6ZpTxyIUphKrusBowYnXP5R/e3FQDzh5EhIwzIvEijJx9ywoeF9EC36Tnz62CRUIPbfla09oVlwGA1wIDAQAB'
);
INSERT INTO production.profiles (id, user_first_name, user_middle_name, user_last_name, user_position, user_company, user_private_key, user_public_key) 
VALUES (4, 'Иван', 'Иванович', 'Иванов', 'Администратор', 'ПАО МОЭСК', 
'MIIEpAIBAAKCAQEAsGSYKiaHzXTqMcIIdGE5iDU/3dlEVNMeUogwFSu83DN9JVNsmlIyhRVmEnm55sB6bs4ES0Lph28EeSmwf71sLWmmtWCxp90Cq3gBzDTzQJqyPNfL3eeFeeNDgNtUs3osB4gMdWs+CvbphRbOsAq3sbNQJ00aGZ6p6zkX/IZxrBIjxY+kxWIjQLy0p/Yn4rybM4+VFKMEUCiXM7Deu16N5hz1FF8HtSxsQBc4gszGWvIaVhZ1iK3y7DRqQBpjHCaA6cKC3Nh6ZpTxyIUphKrusBowYnXP5R/e3FQDzh5EhIwzIvEijJx9ywoeF9EC36Tnz62CRUIPbfla09oVlwGA1wIDAQABAoIBAQCDmons6NJpd9FDToEAU4mZFiGQY4mXv+vfp7w4D2nY4JF+R7+/Y5RNtqlxH2CTyQePpCWQAVw6r5mmzHPi2nDbcPfwWzQxCbP0OpUcxmS2zrQssNRpu1LanbS/buTDA2PWOqsQ7/JaO93+bgXHUje7XQ1wRRY0Byy/UtmSjrxApAqMiFGM4koIu9aOl6CvSb2bxO4MwitV9iS1lMFl4SlaYf9zPtIOJYUQWziRNz1ZK0CKCVnashspM+eNgVw+Z6P5OTOfOPDWBL0ORMvRXLQtoYND8tO9oCS+WV7MaeOk1OP8mP8okrb1o9x9ZGhWIQx7Ue7iaYvzgrUuj0gRtLABAoGBAN+PrpHnaY5DIgGqwBBTzNCaXX1W7n0wALU8qin5jo0cScSttMcM9OLM86BRMSCqYfOPDM0Lak3ZfgybxPCivhwwapsJ7ScP43/BzIFhzIZ2Xtg4P4KrF0OJ/ggfOUV6wgv0OKTzgTQ855lZy3MSKmVb7folXf9obYk+93PzZChvAoGBAMn80SrhuHZF81THBZfbcpJxpKNgTQerX4ia6e/CTT4hjlza8kRGawnWfQOFdmCT3mzdb6tQ+KPsxQWCRtH5DWg84GH/6+lgIUMEB+8OjUoDq6y+XmMWlCePVH1+7swHe7LRlaonwagA7SVn5t8+r5sZc+sIpH1wsc35DOZ+GlIZAoGAd3tBH3WAcqnqeN2bPJ6s7igyIxTc7UdEeZhskXZw+3XM7zKvVVrVXomPA3WhPgYRx6wCeWvKasT8mxx9SuaPmF0//JB3kNLrEZKwC84LEyocUo7tUpbCHjSX8htN7pZHM0BZLb9+pD6QwOK+20cwJW/WZkSmUiSrthhTBENmmj0CgYEAlhIUhju2hYlrRO2ppi4RbeSpYglGshANpr0SWmSOZz8fOrYhkcCP/nsx3s/mJ9M1SsUrFqnOUlyz9WfZnl/gKjYwsB8o8/fMPrJcAq1ZJEid4HaAQjagVNQU/ji0yzo0GaPGAuoO4/fsOgJ8chls91tt2I5PSDPWpyYHA6llfOECgYAhU0FLx2eHwQpFkQyGjgAvnRIzDwB6lKtFwO2BULwu8RBwYgQg1cAr8BCESJkpAHYC4FXEN1PT4qXb3fohGWOoIDWB6LliIUiGEFW5/1rFdLVUnXjA0NCqPQngBE/pW8/tdRnfqkjL07BRbz0RyLa5Me/AZfURkl21flBhvvs7bA==',
'MIIBCgKCAQEAsGSYKiaHzXTqMcIIdGE5iDU/3dlEVNMeUogwFSu83DN9JVNsmlIyhRVmEnm55sB6bs4ES0Lph28EeSmwf71sLWmmtWCxp90Cq3gBzDTzQJqyPNfL3eeFeeNDgNtUs3osB4gMdWs+CvbphRbOsAq3sbNQJ00aGZ6p6zkX/IZxrBIjxY+kxWIjQLy0p/Yn4rybM4+VFKMEUCiXM7Deu16N5hz1FF8HtSxsQBc4gszGWvIaVhZ1iK3y7DRqQBpjHCaA6cKC3Nh6ZpTxyIUphKrusBowYnXP5R/e3FQDzh5EhIwzIvEijJx9ywoeF9EC36Tnz62CRUIPbfla09oVlwGA1wIDAQAB'
);

CREATE TABLE IF NOT EXISTS production.acts (
    id serial PRIMARY KEY,
    staff_id INTEGER NOT NULL,
    superviser_id INTEGER NOT NULL,
    object_id INTEGER NOT NULL,
    review_id INTEGER NOT NULL,
    finished boolean,
    end_at timestamp with time zone,
    staff_sign character varying(2048) DEFAULT '', 
    superviser_sign character varying(2048) DEFAULT '',
    approved boolean,
    reverted boolean,
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

INSERT INTO production.acts (staff_id, superviser_id, object_id, review_id) VALUES (2, 1, 1, 1); 

CREATE TABLE IF NOT EXISTS production.reviews (
    id serial PRIMARY KEY,
    description character varying(2048) DEFAULT '',
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

INSERT INTO production.reviews (description) VALUES ('Осмотр всей ВЛ электромонтёрами');

CREATE TABLE IF NOT EXISTS production.element_types (
    id serial PRIMARY KEY,
    description character varying(2048) DEFAULT '',
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

INSERT INTO production.element_types (description) VALUES ('Опора');
INSERT INTO production.element_types (description) VALUES ('Пролёт');

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

INSERT INTO production.acts_details (act_id, element_id, defects) VALUES (1, 7, '{"defects": [{"defect_id": 1}]}');
INSERT INTO production.acts_details (act_id, element_id, defects) VALUES (1, 39, '{"defects": [{"defect_id": 2}]}');
INSERT INTO production.acts_details (act_id, element_id, defects) VALUES (15, 7, '{"defects": [{"defect_id": 3}]}');
INSERT INTO production.acts_details (act_id, element_id, defects) VALUES (20, 7, '{"defects": [{"defect_id": 4}]}');
INSERT INTO production.acts_details (act_id, element_id, defects) VALUES (50, 7, '{"defects": [{"defect_id": 5}]}');

CREATE TABLE IF NOT EXISTS production.defects (
    id serial PRIMARY KEY,
    element_type INTEGER NOT NULL,
    description character varying(2048) DEFAULT '',
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

INSERT INTO production.defects (element_type, description) VALUES (1, 'Отсутствие диспетчерских названий');
INSERT INTO production.defects (element_type, description) VALUES (1, 'Наличие кустарника под проводами');
INSERT INTO production.defects (element_type, description) VALUES (1, 'Отсутствие плаката, знака безопасности');
INSERT INTO production.defects (element_type, description) VALUES (1, 'Наклон опоры');
INSERT INTO production.defects (element_type, description) VALUES (1, 'Негабарит провода');

CREATE TABLE IF NOT EXISTS production.orders (
    id serial PRIMARY KEY,
    object_id INTEGER NOT NULL,
    tech_tasks jsonb,
    superviser_id INTEGER NOT NULL,
    staff_id INTEGER NOT NULL,
    start_at timestamp with time zone,
    end_at timestamp with time zone,
    superviser_sign character varying(2048) DEFAULT '',
    staff_sign character varying(2048) DEFAULT '',
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

INSERT INTO production.orders (object_id, tech_tasks, superviser_id, staff_id) VALUES (1, '{"tasks": [{ "task_id": 1}, { "task_id": 2}]}', 1, 2);

CREATE TABLE IF NOT EXISTS production.tech_tasks (
    id serial PRIMARY KEY,
    description character varying(2048) DEFAULT '',
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

INSERT INTO production.tech_tasks (description) VALUES ('Не требуется');
INSERT INTO production.tech_tasks (description) VALUES ('Находится в работе под рабочим напряжением');

CREATE TABLE IF NOT EXISTS production.objects (
    id serial PRIMARY KEY,
    object_address character varying(2048) DEFAULT '',
    object_name character varying(2048) DEFAULT '',
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

INSERT INTO production.objects (object_address, object_name) VALUES ('', 'ВЛ 10кВ ТП №265-ТП №240');

CREATE TABLE IF NOT EXISTS production.objects_details (
    object_id INTEGER NOT NULL,
    element_id INTEGER NOT NULL,
    element_name character varying(2048) DEFAULT '',
    element_type INTEGER NOT NULL,
    meta jsonb,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now(),
    deleted_at timestamp with time zone
);

INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 1, 'Опора №1', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 2, 'Опора №2', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 3, 'Опора №3', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 4, 'Опора №4', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 5, 'Опора №5', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 6, 'Опора №6', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 7, 'Опора №7', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 8, 'Опора №8', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 9, 'Опора №9', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 10, 'Опора №10', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 11, 'Опора №11', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 12, 'Опора №12', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 13, 'Опора №13', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 14, 'Опора №14', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 15, 'Опора №15', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 16, 'Опора №16', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 17, 'Опора №17', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 18, 'Опора №18', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 19, 'Опора №19', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 20, 'Опора №20', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 21, 'Опора №21', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 22, 'Опора №22', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 23, 'Опора №23', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 24, 'Опора №24', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 25, 'Опора №25', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 27, 'Опора №26', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 27, 'Опора №27', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 29, 'Пролёт: Опора №1 - Опора №2', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 30, 'Пролёт: Опора №2 - Опора №3', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 31, 'Пролёт: Опора №3 - Опора №4', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 32, 'Пролёт: Опора №4 - Опора №5', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 33, 'Пролёт: Опора №5 - Опора №6', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 34, 'Пролёт: Опора №6 - Опора №7', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 35, 'Пролёт: Опора №7 - Опора №8', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 36, 'Пролёт: Опора №8 - Опора №9', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 37, 'Пролёт: Опора №9 - Опора №10', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 38, 'Пролёт: Опора №10 - Опора №11', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 39, 'Пролёт: Опора №11 - Опора №12', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 40, 'Пролёт: Опора №12 - Опора №13', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 41, 'Пролёт: Опора №13 - Опора №14', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 42, 'Пролёт: Опора №14 - Опора №15', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 43, 'Пролёт: Опора №15 - Опора №16', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 44, 'Пролёт: Опора №16 - Опора №17', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 45, 'Пролёт: Опора №17 - Опора №18', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 46, 'Пролёт: Опора №18 - Опора №19', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 47, 'Пролёт: Опора №19 - Опора №20', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 48, 'Пролёт: Опора №20 - Опора №21', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 49, 'Пролёт: Опора №21 - Опора №22', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 50, 'Пролёт: Опора №22 - Опора №23', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 51, 'Пролёт: Опора №23 - Опора №24', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 52, 'Пролёт: Опора №24 - Опора №25', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 53, 'Пролёт: Опора №25 - Опора №26', 1);
INSERT INTO production.objects_details (object_id, element_id, element_name, element_type) VALUES (1, 54, 'Пролёт: Опора №26 - Опора №27', 1);

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
DROP TABLE production.acts_details;
DROP TABLE production.review;
DROP TABLE production.orders;
DROP TABLE production.objects;
DROP TABLE production.sessions;
DROP TABLE production.objects_details;
DROP TABLE production.defects;