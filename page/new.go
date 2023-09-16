package page

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-shiori/go-readability"
)

type Config struct {
	Client    *http.Client
	Referer   string
	UserAgent string
}

type Option = func(*Config)

func WithClient(client *http.Client) Option {
	return func(c *Config) {
		c.Client = client
	}
}

func WithReferer(referer string) Option {
	return func(c *Config) {
		c.Referer = referer
	}
}

func WithUserAgent(userAgent string) Option {
	return func(c *Config) {
		c.UserAgent = userAgent
	}
}

func New(ctx context.Context, URL *url.URL, options ...Option) (*Page, error) {
	config := &Config{
		Client:    http.DefaultClient,
		Referer:   "",
		UserAgent: "",
	}

	for _, option := range options {
		option(config)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		URL.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Referer", config.Referer)

	res, err := config.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, fmt.Errorf("unsuccessful status: %s", res.Status)
	}

	article, err := readability.FromReader(res.Body, URL)
	if err != nil {
		return nil, err
	}

	document := goquery.NewDocumentFromNode(article.Node)

	mdConverter := md.NewConverter(URL.String(), true, &md.Options{
		EscapeMode: "disabled",
	})
	markdown := mdConverter.Convert(document.Selection)

	selection := document.Find("a")

	anchors := make([]Anchor, 0, selection.Length())
	selection.Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			return
		}

		if strings.HasPrefix(href, "#") {
			return
		}

		anchorURL, err := url.Parse(href)
		if err != nil {
			return
		}

		anchorText := strings.TrimSpace(s.Text())

		anchor := Anchor{
			URL:  anchorURL,
			Text: anchorText,
		}

		anchors = append(anchors, anchor)
	})

	var byline string
	if article.Byline != "__" {
		byline = article.Byline
	}

	return &Page{
		URL:     URL,
		Anchors: anchors,
		Meta: Meta{
			Title:    article.Title,
			Byline:   byline,
			Duration: estimateReadingDuration(article.TextContent),
		},
		Content: Content{
			HTML:     article.Content,
			Markdown: markdown,
		},
	}, nil
}
