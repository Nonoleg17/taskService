CREATE TABLE IF NOT EXISTS "users" (
    id bigserial PRiMARY KEY,
    username varchar NOT NULL,
    email varchar NOT NULL,
    password varchar NOT NULL
    );