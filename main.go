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
	"github.com/lionpuro/portfolio/locales"
	"golang.org/x/text/language"
)

func main() {
	port := os.Getenv("PORT")
	if err := ctxi18n.Load(locales.Content); err != nil {
		log.Fatalf("error loading locales: %v", err)
	}
	mux := http.NewServeMux()

	pageMux := http.NewServeMux()
	pageMux.HandleFunc("GET /", indexHandler)
	pageMux.HandleFunc("GET /blog", blogHandler)
	pageMux.HandleFunc("GET /blog/{slug}", postHandler)

	mux.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", languageMiddleware(pageMux))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: gzipHandler(mux),
	}

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(server.ListenAndServe())
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
