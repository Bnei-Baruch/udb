package models

import (
	"time"
)

type User struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	Email     string    `boil:"email" json:"email" toml:"email" yaml:"email"`
	Name      string    `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	AccountID string    `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
}
