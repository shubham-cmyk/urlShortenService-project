package main

type Request struct {
	URL_LONG  string `json:"url_long"`
	URL_SHORT string `json:"url_short"`
}

type Response struct {
	URL_LONG  string `json:"url_long"`
	URL_SHORT string `json:"url_short"`
}
