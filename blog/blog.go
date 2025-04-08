package blog

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"sort"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark-meta"
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

		metadata, err := parseMetadata(ctx)
		if err != nil {
			return nil, fmt.Errorf("error parsing post date: %v", err)
		}

		post := &Post{
			Slug:    metadata.Slug,
			Title:   metadata.Title,
			Summary: metadata.Summary,
			Date:    metadata.Date,
			Content: buf.String(),
		}

		posts.BySlug[post.Slug] = post
		posts.All = append(posts.All, post)
	}

	sort.SliceStable(posts.All, func(i, j int) bool {
		return posts.All[i].Date.After(posts.All[j].Date)
	})

	return posts, nil
}

func parseMetadata(ctx parser.Context) (PostMetadata, error) {
	metadata := meta.Get(ctx)

	slug := ""
	title := ""
	summary := ""
	date := ""

	if sl, ok := metadata["Slug"].(string); ok {
		slug = sl
	}
	if t, ok := metadata["Title"].(string); ok {
		title = t
	}
	if s, ok := metadata["Summary"].(string); ok {
		summary = s
	}
	if d, ok := metadata["Date"].(string); ok {
		date = d
	}

	if title == "" || summary == "" || date == "" || slug == "" {
		return PostMetadata{}, fmt.Errorf("missing metadata in file %s", slug)
	}

	d, err := parseDate(date)
	if err != nil {
		return PostMetadata{}, fmt.Errorf("error parsing post date: %v", err)
	}

	m := PostMetadata{
		Slug:    slug,
		Title:   title,
		Summary: summary,
		Date:    d,
	}
	return m, nil
}

func parseDate(d string) (time.Time, error) {
	t, err := time.Parse(time.DateOnly, d)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
