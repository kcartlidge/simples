package simples

import (
	"os"
	"strconv"
	"strings"
)

// Config ... Provider of configuration settings.
type Config interface {
	Get(key string, defaultValue string) string
	GetNumber(key string, defaultValue int) int
	SetAllowEnvironmentOverrides(allow bool)
}

type config struct {
	// allowEnvironmentOverrides ... Set to false (default is true) to disable environment overrides.
	allowEnvironmentOverrides bool
	Settings                  map[string]string
}

// CreateConfig ... Initialises the config, loading from a file if it exists.
func CreateConfig(filename string) (Config, error) {
	c := &config{
		allowEnvironmentOverrides: true,
	}
	s, err := loadKeyValues(filename)
	if err == nil {
		c.Settings = s
	}
	return c, err
}

// Get ... Returns a matching value, or the defaultValue if not found.
func (c *config) Get(key string, defaultValue string) string {
	k := strings.ToUpper(key)

	// Check the environment variables.
	if c.allowEnvironmentOverrides {
		for _, e := range os.Environ() {
			kv := strings.Split(e, "=")
			if strings.ToUpper(kv[0]) == k {
				return kv[1]
			}
		}
	}

	// Check the file values.
	v, ok := c.Settings[k]
	if ok {
		return v
	}
	return defaultValue
}

// GetNumber ... Returns a matching value, or the defaultValue if not found.
func (c *config) GetNumber(key string, defaultValue int) int {
	v := c.Get(key, strconv.Itoa(defaultValue))
	iv, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return iv
}

// SetAllowEnvironmentOverrides ... False (default is true) disables environment variable overrides.
func (c *config) SetAllowEnvironmentOverrides(allow bool) {
	c.allowEnvironmentOverrides = allow
}
