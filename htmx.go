package main

import "net/http"

func isHX(r *http.Request) bool {
	return r.Header.Get("HX-request") == "true"
}
