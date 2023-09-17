package page

import (
	_ "embed"
	"net/url"
	"strings"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/metafates/geminite/bookmark"
)

type Page struct {
	URL     *url.URL
	Content Content
	Meta    Meta
	Anchors []Anchor
}

type Meta struct {
	Title    string
	Byline   string
	Duration time.Duration
}

type Content struct {
	HTML, Markdown string
}

type Anchor struct {
	URL  *url.URL
	Text string
}

func (p *Page) Render(renderer *glamour.TermRenderer) (string, error) {
	out, err := renderer.Render(p.Content.Markdown)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(out), nil
}

func (p *Page) AsBookmark() bookmark.Bookmark {
	return bookmark.Bookmark{
		Title: p.Meta.Title,
		URL:   p.URL,
	}
}
