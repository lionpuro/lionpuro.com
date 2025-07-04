package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lionpuro/lionpuro.com/blog"
)

func main() {
	port := os.Getenv("PORT")

	posts, err := blog.ParsePosts("./content/posts")
	if err != nil {
		log.Fatalf("error parsing posts: %v", err)
	}

	rootRouter := http.NewServeMux()

	pageRouter := http.NewServeMux()

	routes := map[string]http.HandlerFunc{
		"GET /":             indexHandler,
		"GET /projects":     projectsHandler,
		"GET /posts":        blogHandler(posts),
		"GET /posts/{slug}": postHandler(posts),
	}
	registerRoutes(pageRouter, routes)

	rootRouter.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir("assets/public"))))

	rootRouter.HandleFunc("/", removeTrailingSlash(
		commonMiddleware(pageRouter),
	))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: rootRouter,
	}

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func registerRoutes(router *http.ServeMux, routes map[string]http.HandlerFunc) {
	for pattern := range routes {
		router.HandleFunc(pattern, routes[pattern])
	}
}

func removeTrailingSlash(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		path := url.Path
		l := len(path) - 1
		if l > 0 && strings.HasSuffix(path, "/") {
			url := path[:l]
			http.Redirect(w, r, url, http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		ctx := context.WithValue(r.Context(), "path", r.URL.Path)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
