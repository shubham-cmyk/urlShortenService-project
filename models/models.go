package models

import "time"

type Request struct {
	URL_LONG   string        `json:"url_long"`
	URL_SHORT  string        `json:"url_short"`
	Expiration time.Duration `json:"expiration"`
}

type Response struct {
	URL_LONG  string `json:"url_long"`
	URL_SHORT string `json:"url_short"`
}
