---
Slug: trailing-slash-redirect-in-go
Title: Trailing slash redirect in Go
Summary: A quick middleware for redirecting routes with a trailing slash to a route without it.
Date: 2025-04-19
---

# Trailing slash redirect in Go

Just a quick middleware for redirecting routes with a trailing slash to a route without it.
So when a user navigates to `/home/` they will be redirected to `/home`.

```go
package main

import (
	"net/http"
	"strings"
)

func redirectTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		l := len(path) - 1
		if l > 0 && strings.HasSuffix(path, "/") {
			url := path[:l]
			http.Redirect(w, r, url, http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
    mux := http.NewServeMux()

	server := &http.Server{
		Handler: redirectTrailingSlash(mux),
	}
}
```
