package blog

import (
	"bytes"
	"cmp"
	"fmt"
	"os"
	"path"
	"slices"
	"strconv"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

type PostMetadata struct {
	Slug    string
	Title   string
	Summary string
	Date    time.Time
	Order   int
}

type Post struct {
	Title   string
	Summary string
	Date    time.Time
	Content string
}

func parseDate(d string) (time.Time, error) {
	t, err := time.Parse(time.DateOnly, d)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func ListPosts() ([]PostMetadata, error) {
	dir := "./posts"
	markdown := goldmark.New(
		goldmark.WithExtensions(meta.Meta),
	)
	var posts []PostMetadata
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

		posts = append(posts, metadata)
	}

	slices.SortFunc(posts, func(a, b PostMetadata) int {
		return cmp.Compare(a.Order, b.Order)
	})
	slices.Reverse(posts)

	return posts, nil
}

func ParsePost(slug string) (*Post, error) {
	content, err := os.ReadFile(fmt.Sprintf("./posts/%s.md", slug))
	if err != nil {
		return nil, fmt.Errorf("error reading markdown file %v", err)
	}

	markdown := goldmark.New(
		goldmark.WithExtensions(meta.Meta),
	)

	ctx := parser.NewContext()

	var buf bytes.Buffer
	if err := markdown.Convert(content, &buf, parser.WithContext(ctx)); err != nil {
		return nil, err
	}

	metadata, err := parseMetadata(ctx)
	if err != nil {
		return nil, fmt.Errorf("error parsing post date: %v", err)
	}

	post := &Post{
		Title:   metadata.Title,
		Summary: metadata.Summary,
		Date:    metadata.Date,
		Content: buf.String(),
	}

	return post, nil
}

func parseMetadata(ctx parser.Context) (PostMetadata, error) {
	metadata := meta.Get(ctx)

	slug := ""
	title := ""
	summary := ""
	date := ""
	order := 999

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
	if o, ok := metadata["Order"].(string); ok {
		if i, err := strconv.Atoi(o); err == nil {
			order = i
		}
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
		Order:   order,
	}
	return m, nil
}
