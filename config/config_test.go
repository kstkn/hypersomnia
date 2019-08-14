package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvironments(t *testing.T) {
	c := Config{
		Environments: "dev:http://dev.example.com;stage:http://stage.example.com",
	}
	assert.Equal(t, map[string]string{"dev": "http://dev.example.com", "stage": "http://stage.example.com"}, c.GetEnvironments())
}

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
