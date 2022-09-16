package models

import "time"

type Request struct {
	URL_LONG   string        `json:"url"`
	URL_SHORT  string        `json:"short"`
	Expiration time.Duration `json:"expiration"`
}

type Response struct {
	URL_LONG  string `json:"url"`
	URL_SHORT string `json:"short"`
}
