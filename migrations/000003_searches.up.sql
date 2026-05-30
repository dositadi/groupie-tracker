BEGIN;

CREATE TABLE IF NOT EXISTS searches (
    id uuid NOT NULL PRIMARY KEY,
    search text NOT NULL,
    userId uuid NOT NULL,
    version integer NOT NULL DEFAULT 1,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT fk_user FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

COMMIT;

