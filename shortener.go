package url_shortener

import (
	"fmt"
	"net/url"
	"sync"
	"sync/atomic"
)

type UrlShortener struct {
	urls    sync.Map
	counter uint32
}

func (shortener *UrlShortener) Shorten(url string) string {
	position := atomic.AddUint32(&(shortener.counter), 1)
	shortPath := fmt.Sprintf("/%x", position)

	shortener.urls.Store(shortPath, url)

	return fmt.Sprintf("https://%s%s", "domain.com", shortPath)
}

func (shortener *UrlShortener) Resolve(url2resolve string) string {
	components, err := url.Parse(url2resolve)
	if err != nil {
		return ""
	}

	resolvedUrl, ok := shortener.urls.Load(components.Path)
	if ok {
		return resolvedUrl.(string)
	} else {
		return ""
	}
}
