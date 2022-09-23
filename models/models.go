package models

import (
	"time"

	"github.com/google/uuid"
)

type Request struct {
	URL        string        `json:"url"`
	Short      string        `json:"short"`
	Expiration time.Duration `json:"time"`
	ID         int           `json:"id"`
}

type Response struct {
	URL        string        `json:"url"`
	SHORT      string        `json:"short"`
	Expiration time.Duration `json:"time"`
	ID         int           `json:"id"`
}

type Uuid struct {
	UUID uuid.UUID
}
