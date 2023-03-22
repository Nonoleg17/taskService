CREATE TABLE IF NOT EXISTS "users" (
    id bigserial PRiMARY KEY,
    username varchar unique NOT NULL,
    email varchar unique NOT NULL,
    password varchar NOT NULL
    );

CREATE TABLE IF NOT EXISTS "tasks" (
  id bigserial PRIMARY KEY,
  header varchar NOT NULL,
  status varchar,
  description varchar,
  creation TIME NOT NULL
);