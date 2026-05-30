BEGIN;

CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL PRIMARY KEY,
    username text NOT NULL,
    email citext NOT NULL UNIQUE,
    hashed_password bytea NOT NULL,
    version integer NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now()
);

COMMIT;
