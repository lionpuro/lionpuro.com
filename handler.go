package main

import (
	"github.com/lionpuro/lionpuro.com/blog"
	"github.com/lionpuro/lionpuro.com/views"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if p != "/" {
		notFoundHandler(w, r)
		return
	}
	title := "Lion Puro"
	desc := "Full stack developer from Finland. Here I document my journey as a developer."
	component := views.FullPage(views.Home(), title, desc)
	component.Render(r.Context(), w)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	title := "Projects - Lion Puro"
	desc := "Full stack developer"
	component := views.FullPage(views.Projects(), title, desc)
	component.Render(r.Context(), w)
}

func blogHandler(posts *blog.Posts) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tag := r.URL.Query().Get("tag")
		psts := posts.All

		switch {
		case tag == "":
			break
		case posts.ByTag[tag] != nil:
			psts = posts.ByTag[tag]
		default:
			notFoundHandler(w, r)
			return
		}

		title := "Posts - Lion Puro"
		desc := "Thoughts and notes, mostly related to web development."
		component := views.FullPage(views.Blog(psts, posts.Tags, tag), title, desc)
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
		component := views.FullPage(
			views.Post(post),
			post.Title,
			post.Summary,
		)
		component.Render(r.Context(), w)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	msg := "Page not found"
	component := views.ErrorPage(http.StatusNotFound, msg)
	component.Render(r.Context(), w)
}
