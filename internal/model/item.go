package model

import "time"

type Item struct {
	Type         string    `json:"type"`
	LastModified time.Time `json:"last-modified"`
	Album        string    `json:"album"`
	Name         string    `json:"name"`
}
