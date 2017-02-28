package simples

import "testing"
import "os"

const filename = "fixtures/test.env"

func TestCreateConfig(t *testing.T) {
	backupOriginal()
	_, err := CreateConfig(filename)

	restoreOriginal()
	if err != nil {
		t.Error("Expected a config, got ", err.Error())
	}
}

func TestGetKey_NoKey_ReturnsDefault(t *testing.T) {
	backupOriginal()
	c, _ := CreateConfig(filename)
	v := c.Get("", "default")

	restoreOriginal()
	if v != "default" {
		t.Error("Expected default value, got ", v)
	}
}

func TestGetKey_KeyNotFound_ReturnsDefault(t *testing.T) {
	backupOriginal()
	c, _ := CreateConfig(filename)
	v := c.Get("NOT-FOUND", "default")

	restoreOriginal()
	if v != "default" {
		t.Error("Expected default value, got ", v)
	}
}

func TestGetKey_KeyExists_ReturnsValue(t *testing.T) {
	backupOriginal()
	c, _ := CreateConfig(filename)
	v := c.Get("PAGE_SIZE", "default")

	restoreOriginal()
	if v != "10" {
		t.Error("Expected 10, got ", v)
	}
}

func TestGetKey_KeyNotFound_EnvNotFound_ReturnsDefault(t *testing.T) {
	backupOriginal()
	c, _ := CreateConfig(filename)
	v := c.Get("NOT-FOUND", "default")

	restoreOriginal()
	if v != "default" {
		t.Error("Expected default value, got ", v)
	}
}

func TestGetKey_KeyFound_EnvNotFound_ReturnsKey(t *testing.T) {
	backupOriginal()
	c, _ := CreateConfig(filename)
	v := c.Get("PAGE_SIZE", "default")

	restoreOriginal()
	if v != "10" {
		t.Error("Expected 10, got ", v)
	}
}

func TestGetKey_KeyCaseDiffers_ReturnsKey(t *testing.T) {
	backupOriginal()
	c, _ := CreateConfig(filename)
	v := c.Get("page_size", "default")

	restoreOriginal()
	if v != "10" {
		t.Error("Expected 10, got ", v)
	}
}

func TestGetKey_EnvNotUppercase_ReturnsDefault(t *testing.T) {
	backupOriginal()
	os.Setenv("max_length", "20")

	c, _ := CreateConfig(filename)
	v := c.Get("MAX_LENGTH", "default")

	restoreOriginal()
	if v != "default" {
		t.Error("Expected 'default', got ", v)
	}
}

func TestGetKey_EnvUppercase_GetNotUppercase_ReturnsEnv(t *testing.T) {
	backupOriginal()
	os.Setenv("MAX_LENGTH", "20")

	c, _ := CreateConfig(filename)
	v := c.Get("max_length", "default")

	restoreOriginal()
	if v != "20" {
		t.Error("Expected 20, got ", v)
	}
}

func TestGetKey_KeyNotFound_EnvFound_ReturnsEnv(t *testing.T) {
	backupOriginal()
	os.Setenv("NOT_IN_FILE", "20")

	c, _ := CreateConfig(filename)
	v := c.Get("NOT_IN_FILE", "default")

	restoreOriginal()
	if v != "20" {
		t.Error("Expected 20, got ", v)
	}
}

func TestGetKey_KeyFound_EnvFound_ReturnsEnv(t *testing.T) {
	backupOriginal()
	os.Setenv("PAGE_SIZE", "20")

	c, _ := CreateConfig(filename)
	v := c.Get("PAGE_SIZE", "default")

	restoreOriginal()
	if v != "20" {
		t.Error("Expected 20, got ", v)
	}
}

// Support.

var (
	hadNotInFile, hadPageSize       bool
	backupNotInFile, backupPageSize string
)

func backupOriginal() {
	backupNotInFile, hadNotInFile = os.LookupEnv("NOT_IN_FILE")
	backupPageSize, hadPageSize = os.LookupEnv("PAGE_SIZE")
	os.Unsetenv("NOT_IN_FILE")
	os.Unsetenv("PAGE_SIZE")
}

func restoreOriginal() {
	if hadNotInFile {
		os.Setenv("NOT_IN_FILE", backupNotInFile)
	}
	if hadPageSize {
		os.Setenv("PAGE_SIZE", backupPageSize)
	}
}
