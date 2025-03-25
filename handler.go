package main

import (
	"context"
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/invopop/ctxi18n/i18n"
	"github.com/lionpuro/portfolio/blog"
	"github.com/lionpuro/portfolio/views"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if p != "/" && p != "/fi" && p != "/en" {
		notFoundHandler(w, r)
		return
	}
	views.FullPage(views.Home(), "Lion Puro", "A full stack web developer").Render(r.Context(), w)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := blog.ListPosts()
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
	views.FullPage(views.Blog(posts), "Blog - Lion Puro", "A full stack web developer").Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post, err := blog.ParsePost(slug)
	if err != nil {
		notFoundHandler(w, r)
		return
	}
	content := unsafe(post.Content)
	views.FullPage(views.Post(post.Title, post.Date, content), "Lion Puro", "Blog").Render(r.Context(), w)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	msg := i18n.T(r.Context(), "not-found")
	views.FullPage(
		views.ErrorPage(http.StatusNotFound, msg), msg, "",
	).Render(r.Context(), w)
}

func unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, html)
		return err
	})
}
