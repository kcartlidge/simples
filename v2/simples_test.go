package simples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	missingFilename    = "fixtures/missing-file.ini"
	noSectionsFilename = "fixtures/no-sections.ini"
	validFilename      = "fixtures/valid.ini"
)

// CreateConfig

func Test_CreateConfig_WithMissingFile_ReturnsError(t *testing.T) {
	_, err := CreateConfig(missingFilename)
	assert.NotNil(t, err)
}

// GetSections

func Test_GetSections_WithNoSections_UsesDefault(t *testing.T) {
	c, _ := CreateConfig(noSectionsFilename)
	s := c.GetSections()
	if assert.Len(t, s, 1) {
		assert.Equal(t, s[0], "DEFAULT")
	}
}

func Test_GetSections_WithSections_UsesSectionsAndDefault(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetSections()
	if assert.Len(t, s, 4) {
		assert.Contains(t, s, "DEFAULT")
		assert.Contains(t, s, "SECTION 1")
		assert.Contains(t, s, "SECTION 2")
		assert.Contains(t, s, "SECTION 3")
	}
}

// GetSection

func Test_GetSection_WithInvalidSection_ReturnsEmpty(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetSection("MISSING")
	assert.Len(t, s, 0)
}

func Test_GetSection_WithSection_ReturnsSection(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetSection("SECTION 3")
	if assert.Len(t, s, 1) {
		assert.Contains(t, s[1].Section, "SECTION 3")
	}
}

func Test_GetSection_WithSection_ReturnsExpectedSectionSize(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetSection("SECTION 1")
	assert.Len(t, s, 2)
}

func Test_GetSection_WithDuplicatedSection_ReturnsMerged(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetSection("SECTION 2")
	assert.Len(t, s, 6)
}

func Test_GetSection_WithSection_ReturnsSequenced(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetSection("SECTION 2")
	if assert.Len(t, s, 6) {
		assert.Equal(t, 1, s[1].Sequence)
		assert.Equal(t, 2, s[2].Sequence)
		assert.Equal(t, 3, s[3].Sequence)
		assert.Equal(t, 4, s[4].Sequence)
		assert.Equal(t, 5, s[5].Sequence)
		assert.Equal(t, 6, s[6].Sequence)
	}
}

// GetString

func Test_GetString_WithMissingSection_ReturnsDefault(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetString("MISSING-SECTION", "KEY", "default")
	assert.Equal(t, "default", s)
}

func Test_GetString_WithMissingKey_ReturnsDefault(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetString("SECTION 1", "KEY", "default")
	assert.Equal(t, "default", s)
}

func Test_GetString_WithValidSectionAndKey_ReturnsValue(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetString("SECTION 2", "Section Two Example Key 2", "default")
	assert.Equal(t, "Example Value 2 for Section Two", s)
}

func Test_GetString_WithIncorrectSectionCase_ReturnsValue(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetString("Section 2", "Section Two Example Key 2", "default")
	assert.Equal(t, "Example Value 2 for Section Two", s)
}

func Test_GetString_WithIncorrectKeyCase_ReturnsValue(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetString("SECTION 2", "SECTION TWO EXAMPLE KEY 2", "default")
	assert.Equal(t, "Example Value 2 for Section Two", s)
}

// GetNumber

func Test_GetNumber_WithNoValue_ReturnsDefault(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetNumber("DEFAULT", "MissingNumber", 99)
	assert.Equal(t, 99, s)
}

func Test_GetNumber_WithNonNumericValue_ReturnsDefault(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetNumber("SECTION 2", "Section Two Example Key 2", 99)
	assert.Equal(t, 99, s)
}

func Test_GetNumber_WithValue_ReturnsValue(t *testing.T) {
	c, _ := CreateConfig(validFilename)
	s := c.GetNumber("DEFAULT", "ExampleNumber", 99)
	assert.Equal(t, 10, s)
}
