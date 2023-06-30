CREATE TABLE
    "users" (
        "id" bigserial PRIMARY KEY,
        "name" varchar NOT NULL,
        "token" varchar NULL,
        "user_name " varchar NOT NULL,
        "password" varchar NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );