package url_shortener

import (
	"fmt"
	"net/url"
	"sync"
	"sync/atomic"
)

var urls sync.Map
var counter uint32

func Shorten(url string) string {
	position := atomic.AddUint32(&counter, 1)
	shortPath := fmt.Sprintf("/%x", position)

	urls.Store(shortPath, url)

	return fmt.Sprintf("https://%s%s", "domain.com", shortPath)
}

func Resolve(url2resolve string) string {
	components, err := url.Parse(url2resolve)
	if err != nil {
		return ""
	}

	resolvedUrl, ok := urls.Load(components.Path)
	if ok {
		return resolvedUrl.(string)
	} else {
		return ""
	}
}

func clear() {
	urls = sync.Map{}
	counter = 0
}
