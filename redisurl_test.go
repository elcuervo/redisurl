package redisurl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Parse(t *testing.T) {
	var url *Url

	url = Parse("redis://:6379/")
	assert.Equal(t, 6379, url.Port)
	assert.Equal(t, "127.0.0.1", url.Host)
	assert.Equal(t, 0, url.Database)

	url = Parse("redis://username:passwd@somewhere:6380/3")
	assert.Equal(t, 6380, url.Port)
	assert.Equal(t, 3, url.Database)
	assert.Equal(t, "somewhere", url.Host)
	assert.Equal(t, "passwd", url.Password)

	url = Parse("redis://:@somewhere:6380/9")
	assert.Equal(t, 9, url.Database)
	assert.Empty(t, url.Password)
}
