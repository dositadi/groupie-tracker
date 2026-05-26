package groupietracker

import (
	"embed"
)

//go:embed migrations
var embeddedFiles embed.FS

type Migrations struct {
	embedded embed.FS
}

func NewMigrations() *Migrations {
	return &Migrations{
		embedded: embeddedFiles,
	}
}

func (m *Migrations) Get() embed.FS {
	return m.embedded
}
