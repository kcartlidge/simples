package simples

import (
	"os"
	"testing"
)

const (
	filename = "fixtures/test.env"

	// Unique enough to not exist in the os.env
	notInFile = "SIMPLES_CONFIG_TEST_NOT_IN_FILE"
	pageSize  = "SIMPLES_CONFIG_TEST_PAGE_SIZE"
)

func setup() func() {
	os.Unsetenv(notInFile)
	os.Unsetenv(pageSize)
	return func() {}
}

func Test_CreateConfig(t *testing.T) {
	defer setup()()
	_, err := CreateConfig(filename)

	if err != nil {
		t.Error("Expected a config, got ", err.Error())
	}
}

func Test_NoKey_ReturnsDefault(t *testing.T) {
	defer setup()()
	c, _ := CreateConfig(filename)
	v := c.Get("", "default")

	if v != "default" {
		t.Error("Expected default value, got ", v)
	}
}

func Test_KeyNotFound_ReturnsDefault(t *testing.T) {
	defer setup()()
	c, _ := CreateConfig(filename)
	v := c.Get("NOT-FOUND", "default")

	if v != "default" {
		t.Error("Expected default value, got ", v)
	}
}

func Test_ExistsInFile_And_NotInEnv_ReturnsFromFile(t *testing.T) {
	defer setup()()
	c, _ := CreateConfig(filename)
	v := c.Get(pageSize, "default")

	if v != "10" {
		t.Error("Expected 10, got ", v)
	}
}

func Test_NotExistsInFile_And_ExistsInEnv_ReturnsFromEnv(t *testing.T) {
	defer setup()()
	os.Setenv(notInFile, "20")

	c, _ := CreateConfig(filename)
	v := c.Get(notInFile, "default")

	if v != "20" {
		t.Error("Expected 20, got ", v)
	}
}

func Test_ExistsInFile_And_ExistsInEnv_ReturnsFromEnv(t *testing.T) {
	defer setup()()
	os.Setenv(pageSize, "20")

	c, _ := CreateConfig(filename)
	v := c.Get(pageSize, "default")

	if v != "20" {
		t.Error("Expected 20, got ", v)
	}
}

func Test_CaseDiffersInFile_AndNotInEnv_StillReturnsFromFile(t *testing.T) {
	defer setup()()
	c, _ := CreateConfig(filename)
	v := c.Get(pageSize, "default")

	if v != "10" {
		t.Error("Expected 10, got ", v)
	}
}

func Test_CaseDiffersInEnv_StillReturnsFromEnv(t *testing.T) {
	defer setup()()
	os.Setenv("max_length", "20")

	c, _ := CreateConfig(filename)
	v := c.Get("MAX_LENGTH", "default")

	if v != "20" {
		t.Error("Expected '20', got ", v)
	}
}

func Test_GetNumber_DoesNotExist_ReturnsDefault(t *testing.T) {
	defer setup()()
	c, _ := CreateConfig(filename)
	v := c.GetNumber("NOT_FOUND", 5)

	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}

func Test_GetNumber_Exists_ReturnsValue(t *testing.T) {
	defer setup()()
	c, _ := CreateConfig(filename)
	v := c.GetNumber(pageSize, 5)

	if v != 10 {
		t.Error("Expected 10, got ", v)
	}
}
