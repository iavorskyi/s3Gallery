package model

import (
	"testing"
	"time"
)

func TestUser(t *testing.T) *User {
	return &User{
		Email:     "user@example.com",
		Password:  "12345",
		FirstName: "Jack",
		LastName:  "Black",
	}
}

func TestItem(t *testing.T) *Item {
	return &Item{
		Type:         "jpeg",
		LastModified: time.Now(),
		Album:        "test",
		Name:         "test-item",
	}
}
