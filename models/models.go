package models

import "time"

type Request struct {
	URL        string        `json:"url"`
	Short      string        `json:"short"`
	Expiration time.Duration `json:"time"`
}

type Response struct {
	URL        string        `json:"url"`
	SHORT      string        `json:"short"`
	Expiration time.Duration `json:"time"`
}
