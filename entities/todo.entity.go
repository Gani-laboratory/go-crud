package entities

import "gorm.io/gorm"

// schema dari tabel Todo
// kita coba dengan jika datanya null
// jika return datanya ada yg null, silahkan pake NullString, contohnya dibawah
// Penulis       config.NullString `json:"penulis"`
type Todo struct {
	gorm.Model
	Title    	string `json:"title"`
	Penulis		string `json:"penulis"`
}