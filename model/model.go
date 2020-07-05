package model

import (
	"gioui.org/widget"
	"net"
)

type Client struct {
	Socket net.Conn
	data   chan []byte
}

type WingCalEcommands map[int]WingCalEcommand

type WingCalEcommand struct {
	Id       string
	Type     string
	Name     string      `json:"name"`
	Enabled  bool        `json:"enabled"`
	CompType string      `json:"comptype"`
	SubType  string      `json:"subtype"`
	Command  interface{} `json:"command"`
	Data     interface{} `json:"data"`
}
type EditabilnaPoljaVrsteRadova struct {
	Id        *widget.Editor
	Naziv     *widget.Editor
	Opis      *widget.Editor
	Obracun   *widget.Editor
	Jedinica  *widget.Editor
	Cena      *widget.Editor
	Slug      *widget.Editor
	Omogucen  *widget.Bool
	Materijal map[int]*widget.Editor
}

type ElementMenu struct {
	ID        string
	Title     string
	Materijal bool
}
