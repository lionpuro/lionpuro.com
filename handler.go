package main

import (
	"context"
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/invopop/ctxi18n/i18n"
	"github.com/lionpuro/lionpuro.com/blog"
	"github.com/lionpuro/lionpuro.com/views"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if p != "/" && p != "/fi" && p != "/en" {
		notFoundHandler(w, r)
		return
	}
	title := i18n.T(r.Context(), "meta.home.title")
	desc := i18n.T(r.Context(), "meta.home.description")
	component := views.FullPage(views.Home(false), title, desc)
	if isHX(r) {
		component = views.Home(true)
	}
	component.Render(r.Context(), w)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	title := i18n.T(r.Context(), "meta.projects.title")
	desc := i18n.T(r.Context(), "meta.projects.description")
	component := views.FullPage(views.Projects(false), title, desc)
	if isHX(r) {
		component = views.Projects(true)
	}
	component.Render(r.Context(), w)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := blog.ListPosts()
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
	title := i18n.T(r.Context(), "meta.blog.title")
	desc := i18n.T(r.Context(), "meta.blog.description")
	component := views.FullPage(views.Blog(false, posts), title, desc)
	if isHX(r) {
		component = views.Blog(true, posts)
	}
	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post, err := blog.ParsePost(slug)
	if err != nil {
		notFoundHandler(w, r)
		return
	}
	content := unsafe(post.Content)
	component := views.FullPage(
		views.Post(false, post.Title, post.Date, content),
		post.Title,
		post.Summary,
	)
	if isHX(r) {
		component = views.Post(true, post.Title, post.Date, content)
	}
	component.Render(r.Context(), w)
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
