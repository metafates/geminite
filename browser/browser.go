package browser

import (
	"context"
	"net/http"
	"net/url"

	"github.com/metafates/geminite/cache"
	"github.com/metafates/geminite/page"
)

type Browser struct {
	cache     *cache.Cache[string, *page.Page]
	client    *http.Client
	userAgent string
}

func New() (*Browser, error) {
	c, err := cache.New[string, *page.Page]()
	if err != nil {
		return nil, err
	}

	return &Browser{
		userAgent: "geminite",
		cache:     c,
		client:    http.DefaultClient,
	}, nil
}

func (b *Browser) Open(ctx context.Context, URL *url.URL) (*page.Page, error) {
	p, found, err := b.cache.Get(URL.String())
	if err != nil {
		return nil, err
	}

	if found {
		return p, nil
	}

	p, err = page.New(
		ctx,
		URL,
		page.WithUserAgent(b.userAgent),
		page.WithClient(b.client),
	)
	if err != nil {
		return nil, err
	}

	if err = b.cache.Set(URL.String(), p); err != nil {
		return nil, err
	}

	return p, nil
}
