package groupietracker

import (
	"embed"
)

//go:embed migrations src/output.css internal/web/static/auth/* internal/web/static/partials/auth/* internal/web/static/pages/*
//go:embed internal/web/static/partials/pages/*
var embeddedFiles embed.FS

type Embedded struct {
	embedded embed.FS
}

func New() *Embedded {
	return &Embedded{
		embedded: embeddedFiles,
	}
}

func (m *Embedded) Get() embed.FS {
	return m.embedded
}

func (m *Embedded) GetPath(dir, fileName string) string {
	switch dir {
	case "auth":
		return "internal/web/static/auth/" + fileName
	}
	return ""
}
