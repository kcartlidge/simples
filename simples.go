package simples

import "os"
import "strings"

// Config ... Provider of configuration settings.
type Config interface {
	Get(key string, defaultValue string) string
}

type config struct {
	Settings map[string]string
}

// CreateConfig ... Initialises the config, loading from a file if it exists.
func CreateConfig(filename string) (Config, error) {
	c := config{}
	s, err := loadKeyValues(filename)
	if err == nil {
		c.Settings = s
	}
	return c, err
}

// Get ... Returns a matching value, or the defaultValue if not found.
func (c config) Get(key string, defaultValue string) string {
	k := strings.ToUpper(key)

	// Check the environment variables.
	v, ok := os.LookupEnv(k)
	if ok {
		return v
	}

	// Check the file values.
	v, ok = c.Settings[k]
	if ok {
		return v
	}
	return defaultValue
}
