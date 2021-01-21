package config_test

import (
	"os"
	"testing"

	"github.com/barismar/planner-rest-api/config"
	"github.com/stretchr/testify/assert"
)

func Test_EnvHasBeenSetBefore_EnvShouldMatchWithValueHasBeenSet(t *testing.T) {
	os.Setenv("KEY", "value")

	assert.Equal(t, os.Getenv("KEY"), "value")
	assert.Equal(t, "value", config.Env("KEY", "default_value"))

	os.Unsetenv("KEY")
}

func Test_EnvHasNotBeenSetBefore_EnvShouldMatchWithDefaultValue(t *testing.T) {
	assert.Equal(t, os.Getenv("KEY"), "")
	assert.Equal(t, "default_value", config.Env("KEY", "default_value"))
}
