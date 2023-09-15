package page

import (
	"bytes"
	"context"
	_ "embed"
	"net/url"
	"strings"
	"text/template"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/charmbracelet/glamour"
)

//go:embed page.md.tmpl
var markdownTemplateFile []byte

var markdownTemplate = template.Must(template.New("markdown").Parse(string(markdownTemplateFile)))

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

func (p *Page) markdownTemplate() (string, error) {
	buffer := new(bytes.Buffer)

	if err := markdownTemplate.Execute(buffer, p); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func (p *Page) Render(renderer *glamour.TermRenderer) (string, error) {
	tmpl, err := p.markdownTemplate()
	if err != nil {
		return "", err
	}

	out, err := renderer.Render(tmpl)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(out), nil
}

func (p *Page) Descendant(ctx context.Context, URL *url.URL) (*Page, error) {
	return New(ctx, URL, WithReferer(p.URL.String()))
}
