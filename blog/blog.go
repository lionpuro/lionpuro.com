package blog

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"sort"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type Posts struct {
	All    []*Post
	BySlug map[string]*Post
}

type PostMetadata struct {
	Slug    string
	Title   string
	Summary string
	Date    time.Time
}

type Post struct {
	Slug    string
	Title   string
	Summary string
	Date    time.Time
	Content string
}

func ParsePosts(dir string) (*Posts, error) {
	posts := &Posts{
		All:    []*Post{},
		BySlug: make(map[string]*Post),
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
			highlighting.NewHighlighting(
				highlighting.WithCustomStyle(customTheme()),
			),
			extension.Table,
		),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("error reading markdown posts %v", err)
	}

	for _, file := range files {
		contents, err := os.ReadFile(path.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		ctx := parser.NewContext()

		var buf bytes.Buffer
		if err := markdown.Convert(contents, &buf, parser.WithContext(ctx)); err != nil {
			return nil, err
		}

		metadata := meta.Get(ctx)

		post, err := newPost(metadata, buf.String())
		if err != nil {
			return nil, fmt.Errorf("new post: %v", err)
		}

		posts.BySlug[post.Slug] = post
		posts.All = append(posts.All, post)
	}

	sort.SliceStable(posts.All, func(i, j int) bool {
		return posts.All[i].Date.After(posts.All[j].Date)
	})

	return posts, nil
}

func newPost(metadata map[string]interface{}, content string) (*Post, error) {
	slug, ok := metadata["Slug"].(string)
	if !ok {
		return nil, errors.New("invalid slug")
	}
	title, ok := metadata["Title"].(string)
	if !ok {
		return nil, errors.New("invalid title")
	}
	summary, ok := metadata["Summary"].(string)
	if !ok {
		return nil, errors.New("invalid summary")
	}
	dateString, ok := metadata["Date"].(string)
	if !ok {
		return nil, errors.New("invalid date")
	}
	date, err := parseDate(dateString)
	if err != nil {
		return nil, errors.New("invalid date")
	}

	post := &Post{
		Slug:    slug,
		Title:   title,
		Summary: summary,
		Date:    date,
		Content: content,
	}
	return post, nil
}

func parseDate(d string) (time.Time, error) {
	t, err := time.Parse(time.DateOnly, d)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
