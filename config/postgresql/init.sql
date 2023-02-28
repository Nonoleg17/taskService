CREATE TABLE IF NOT EXISTS "users" (
    id bigserial PRiMARY KEY,
    username varchar unique NOT NULL,
    email varchar NOT NULL,
    password varchar NOT NULL
    );