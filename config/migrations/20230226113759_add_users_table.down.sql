CREATE TABLE "users" {
    "id" bigserial PROMARY KEY
    "username" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
}