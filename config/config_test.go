package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var registryTestData = []struct {
	in  string
	out string
}{
	{"", "mdns"},
	{"mdns", "mdns"},
	{"consul", "consul"},
}

func TestGetRegistry(t *testing.T) {
	for _, tt := range registryTestData {
		c := Config{
			Registry: tt.in,
		}
		assert.Equal(t, tt.out, c.GetRegistry().String())
	}
}
