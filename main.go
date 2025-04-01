package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/invopop/ctxi18n"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lionpuro/lionpuro.com/locales"
	"golang.org/x/text/language"
)

func main() {
	port := os.Getenv("PORT")
	if err := ctxi18n.Load(locales.Content); err != nil {
		log.Fatalf("error loading locales: %v", err)
	}
	rootRouter := http.NewServeMux()

	pageRouter := http.NewServeMux()

	routes := map[string]http.HandlerFunc{
		"GET /":            indexHandler,
		"GET /projects":    projectsHandler,
		"GET /blog":        blogHandler,
		"GET /blog/{slug}": postHandler,
	}
	registerRoutes(pageRouter, routes)

	rootRouter.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	rootRouter.HandleFunc("/", removeTrailingSlash(
		commonMiddleware(languageMiddleware(pageRouter)),
	))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: gzipHandler(rootRouter),
	}

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func registerRoutes(router *http.ServeMux, routes map[string]http.HandlerFunc) {
	for pattern := range routes {
		router.HandleFunc(pattern, routes[pattern])
		router.HandleFunc(withBase("fi", pattern), routes[pattern])
		router.HandleFunc(withBase("en", pattern), routes[pattern])
	}
}

func withBase(base, pattern string) string {
	parts := strings.Fields(pattern)
	if len(parts) == 1 {
		return path.Join("/", base, parts[0])
	}
	newPath := path.Join("/", base, parts[1])
	return parts[0] + " " + newPath
}

func detectLanguage(r *http.Request) string {
	lang := "en"
	accept := r.Header.Get("Accept-Language")
	matcher := language.NewMatcher([]language.Tag{
		language.English, // fallback
		language.Finnish,
	})
	tag, _ := language.MatchStrings(matcher, accept)
	b, _ := tag.Base()

	if b.String() == "fi" {
		lang = "fi"
	}
	return lang
}

func languageMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathSegments := strings.Split(r.URL.Path, "/")
		langParam := pathSegments[1]
		lang := detectLanguage(r)
		urlBase := "/"

		if lang == "fi" && langParam == "" {
			url := path.Join("/", "fi")
			http.Redirect(w, r, url, http.StatusFound)
			return
		}

		if langParam == "fi" || langParam == "en" {
			lang = langParam
			urlBase = "/" + langParam
		}

		i18nCtx, err := ctxi18n.WithLocale(r.Context(), lang)
		ctx := context.WithValue(i18nCtx, "urlBase", urlBase)
		if err != nil {
			log.Printf("error setting locale: %v", err)
			http.Error(w, "error setting locale", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
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
		next.ServeHTTP(w, r)
	})
}
