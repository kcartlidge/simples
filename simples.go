package simples

import (
	"strconv"
	"strings"
)

// Config ... Provider of configuration settings.
type Config interface {
	GetSections() []string

	// GetSection ... Result sequence matches original file. Iteration over result is random.
	GetSection(section string) map[int]Entry

	GetString(section, key, defaultValue string) string
	GetNumber(section, key string, defaultValue int) int
}

type config struct {
	Settings Sections
}

// CreateConfig ... Initialises the config, loading from a file if it exists.
func CreateConfig(filename string) (Config, error) {
	c := &config{
		Settings: make(Sections),
	}
	s, err := loadKeyValues(filename)
	if err == nil {
		c.Settings = s
	}
	return c, err
}

// GetSections ... Returns the name of all sections found.
func (c *config) GetSections() []string {
	res := []string{}
	for name := range c.Settings {
		res = append(res, name)
	}
	return res
}

// GetSections ... Returns the name of all sections found.
func (c *config) GetSection(section string) map[int]Entry {
	res := make(map[int]Entry)
	su := strings.ToUpper(section)
	if s, ok := c.Settings[su]; ok {
		for _, e := range s {
			res[e.Sequence] = e
		}
	}
	return res
}

// GetString ... Returns the value or the given default.
func (c *config) GetString(section, key, defaultValue string) string {
	su := strings.ToUpper(section)
	if s, ok := c.Settings[su]; ok {
		ku := strings.ToUpper(key)
		for _, e := range s {
			if e.KeyUpper == ku {
				return e.Value
			}
		}
	}
	return defaultValue
}

// GetNumber ... Returns a matching value, or the defaultValue if not found.
func (c *config) GetNumber(section, key string, defaultValue int) int {
	v := c.GetString(section, key, strconv.Itoa(defaultValue))
	iv, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return iv
}
