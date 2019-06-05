package url_shortener

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShorten(t *testing.T) {
	clear()

	assert.Equal(t, "https://domain.com/1", Shorten("http://test.ru/asdfasdf"))
	assert.Equal(t, "https://domain.com/2", Shorten("http://test.ru/fasdf"))

	assert.Equal(t, "http://test.ru/asdfasdf", Resolve("https://domain.com/1"))
	assert.Equal(t, "http://test.ru/fasdf", Resolve("https://domain.com/2"))
}

func TestNonExistent(t *testing.T) {
	clear()

	assert.Equal(t, "", Resolve("invalid url"))
	assert.Equal(t, "", Resolve("https://valid.ru/non-existent"))
}
