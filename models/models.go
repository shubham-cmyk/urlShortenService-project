package main

type request struct {
	URL_LONG  string `json:"url_long"`
	URL_SHORT string `json:"url_short"`
}

type response struct {
	URL_LONG  string `json:"url_long"`
	URL_SHORT string `json:"url_short"`
}
