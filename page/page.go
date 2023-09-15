package page

import (
	"context"
	_ "embed"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/charmbracelet/glamour"
)

type Page struct {
	document *goquery.Document
	URL      *url.URL
	Content  Content
	Meta     Meta
	Anchors  []Anchor
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

func (p *Page) Descendant(ctx context.Context, URL *url.URL) (*Page, error) {
	return New(ctx, URL, WithReferer(p.URL.String()))
}
