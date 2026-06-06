package data

import "time"

/*
CREATE TABLE IF NOT EXISTS preference (
    id uuid NOT NULL PRIMARY KEY,
    userId uuid NOT NULL,
    filter filters NOT NULL DEFAULT 'ID',
    sort sorts NOT NULL DEFAULT 'ASC',
    version integer NOT NULL DEFAULT 1,
    createdAt TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updateAt TIMESTAMP WITH TIME ZONE DEFAULT now(),

    CONSTRAINT fk_user FOREIGN KEY (userId) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
);
*/

type Preference struct {
	Id        string
	UserId    string
	Filter    string
	Sort      string
	Version   string
	CreatedAt time.Time
	UpdateAt  time.Time
}

type PreferenceUpdate struct {
	Id     string
	UserId string
	Filter *string
	Sort   *string
}
