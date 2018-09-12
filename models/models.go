package models

//go:generate kallax gen

import kallax "gopkg.in/src-d/go-kallax.v1"

// ABC ...
type ABC struct {
	kallax.Model `table:"a" pk:"id,autoincr"`
	kallax.Timestamps
	ID   int64
	Name string
}

func newABC(name string) *ABC {
	return &ABC{Name: name}
}
