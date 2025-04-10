package main

import (
	"context"
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/lionpuro/lionpuro.com/blog"
	"github.com/lionpuro/lionpuro.com/views"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if p != "/" {
		notFoundHandler(w, r)
		return
	}
	title := "Lion Puro"
	desc := "Full stack developer from Finland. Here I document my journey as a developer."
	component := views.FullPage(views.Home(false), title, desc)
	if isHX(r) {
		component = views.Home(true)
	}
	component.Render(r.Context(), w)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	title := "Projects - Lion Puro"
	desc := "Full stack developer"
	component := views.FullPage(views.Projects(false), title, desc)
	if isHX(r) {
		component = views.Projects(true)
	}
	component.Render(r.Context(), w)
}

func blogHandler(posts *blog.Posts) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title := "Posts - Lion Puro"
		desc := "Thoughts and notes, mostly related to web development."
		component := views.FullPage(views.Blog(false, posts.All), title, desc)
		if isHX(r) {
			component = views.Blog(true, posts.All)
		}
		component.Render(r.Context(), w)
	}
}

func postHandler(posts *blog.Posts) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		post := posts.BySlug[slug]
		if post == nil {
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
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	msg := "Page not found"
	component := views.ErrorPage(http.StatusNotFound, msg)
	if !isHX(r) {
		component = views.FullPage(component, msg, "")
	}
	component.Render(r.Context(), w)
}

func unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, html)
		return err
	})
}
