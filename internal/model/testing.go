package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:     "user@example.com",
		Password:  "12345",
		FirstName: "Jack",
		LastName:  "Black",
	}
}
