BEGIN;

CREATE TABLE IF NOT EXISTS favorites (
    id uuid NOT NULL PRIMARY KEY,
    userId uuid NOT NULL,
    artistId integer NOT NULL,
    status boolean NOT NULL DEFAULT false,
    version integer NOT NULL DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE  NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),

    CONSTRAINT fk_user FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE EXTENSION pg_trgm;

CREATE INDEX idx_artist_id ON favorites (artistId);

COMMIT;