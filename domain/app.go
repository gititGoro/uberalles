package domain

import (
	"github.com/satori/go.uuid"
)

type app struct{
	name string
	key uuid.UUID
}

func newApp(name string) *app{
	return &app{name,uuid.NewV4()}
}