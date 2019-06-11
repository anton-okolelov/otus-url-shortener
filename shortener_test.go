package url_shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShorten(t *testing.T) {
	shortener := UrlShortener{}
	assert.Equal(t, "https://domain.com/1", shortener.Shorten("http://test.ru/asdfasdf"))
	assert.Equal(t, "https://domain.com/2", shortener.Shorten("http://test.ru/fasdf"))

	assert.Equal(t, "http://test.ru/asdfasdf", shortener.Resolve("https://domain.com/1"))
	assert.Equal(t, "http://test.ru/fasdf", shortener.Resolve("https://domain.com/2"))
}

func TestNonExistent(t *testing.T) {
	shortener := UrlShortener{}
	assert.Equal(t, "", shortener.Resolve("invalid url"))
	assert.Equal(t, "", shortener.Resolve("https://valid.ru/non-existent"))
}
