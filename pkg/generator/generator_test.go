package generator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whinarn/strongpass/pkg/generator"
)

func TestNewWithoutConfigShouldSucceed(t *testing.T) {
	generator, err := generator.New(nil)
	assert.NoError(t, err)
	assert.NotNil(t, generator)
}

func TestDefaultConfigShouldSucceed(t *testing.T) {
	generatorConfig := generator.DefaultConfig()
	assert.NotNil(t, generatorConfig)
}

func TestNewWithDefaultConfigShouldSucceed(t *testing.T) {
	generatorConfig := generator.DefaultConfig()
	generator, err := generator.New(generatorConfig)
	assert.NoError(t, err)
	assert.NotNil(t, generator)
}

func TestNewWithCustomConfigShouldSucceed(t *testing.T) {
	generatorConfig := generator.Config{
		AllowLowerCaseLetters: true,
		AllowUpperCaseLetters: true,
		AllowDigits:           true,
		AllowSpecials:         true,
		MinLength:             10,
		MaxLength:             20,
	}
	generator, err := generator.New(&generatorConfig)
	assert.NoError(t, err)
	assert.NotNil(t, generator)
}

func TestNewWithZeroLengthShouldFail(t *testing.T) {
	generatorConfig := generator.Config{}
	generator, err := generator.New(&generatorConfig)
	assert.Error(t, err)
	assert.Nil(t, generator)
	assert.Contains(t, err.Error(), "minumum length")
	assert.Contains(t, err.Error(), "cannot be zero or negative")
}

func TestNewWithNegativeLengthShouldFail(t *testing.T) {
	generatorConfig := generator.Config{
		MinLength: -20,
		MaxLength: -20,
	}
	generator, err := generator.New(&generatorConfig)
	assert.Error(t, err)
	assert.Nil(t, generator)
	assert.Contains(t, err.Error(), "minumum length")
	assert.Contains(t, err.Error(), "cannot be zero or negative")
}

func TestNewWithLowerMaximumLengthShouldFail(t *testing.T) {
	generatorConfig := generator.Config{
		MinLength: 10,
		MaxLength: 5,
	}
	generator, err := generator.New(&generatorConfig)
	assert.Error(t, err)
	assert.Nil(t, generator)
	assert.Contains(t, err.Error(), "maximum length")
	assert.Contains(t, err.Error(), "cannot be lower")
}

func TestNewWithNegativeMinimumCharacterCountsShouldSucceed(t *testing.T) {
	generatorConfig := generator.Config{
		AllowLowerCaseLetters: true,
		MinLength:             10,
		MaxLength:             10,
		MinLowerCaseLetters:   -1,
		MinUpperCaseLetters:   -1,
		MinDigits:             -1,
		MinSpecials:           -1,
	}
	generator, err := generator.New(&generatorConfig)
	assert.NoError(t, err)
	assert.NotNil(t, generator)
}

func TestNewWithNoCharsShouldFail(t *testing.T) {
	generatorConfig := generator.Config{
		CharSet:               nil,
		AllowLowerCaseLetters: false,
		AllowUpperCaseLetters: false,
		AllowDigits:           false,
		AllowSpecials:         false,
		MinLength:             10,
		MaxLength:             20,
	}
	generator, err := generator.New(&generatorConfig)
	assert.Error(t, err)
	assert.Nil(t, generator)
	assert.Contains(t, err.Error(), "no characters available")
}

func TestNewWithNotEnoughLengthShouldFail(t *testing.T) {
	generatorConfig := generator.Config{
		AllowLowerCaseLetters: true,
		MinLength:             1,
		MaxLength:             1,
		MinLowerCaseLetters:   1,
		MinUpperCaseLetters:   1,
		MinDigits:             1,
		MinSpecials:           1,
	}
	generator, err := generator.New(&generatorConfig)
	assert.Error(t, err)
	assert.Nil(t, generator)
	assert.Contains(t, err.Error(), "minimum length")
}

func TestNewWithCustomCharSetShouldSucceed(t *testing.T) {
	generatorConfig := generator.Config{
		CharSet:   []rune("0123456789abcdef"),
		MinLength: 10,
		MaxLength: 20,
	}
	generator, err := generator.New(&generatorConfig)
	assert.NoError(t, err)
	assert.NotNil(t, generator)
}

func TestGeneratorGeneratePasswordDefaultShouldSucceed(t *testing.T) {
	generator, _ := generator.New(nil)
	password := generator.GeneratePassword()
	assert.NotEmpty(t, password)
}

func TestGeneratorGeneratePasswordWithLengthShouldSucceed(t *testing.T) {
	generatorConfig := generator.Config{
		AllowLowerCaseLetters: true,
		MinLength:             10,
		MaxLength:             10,
	}
	generator, _ := generator.New(&generatorConfig)
	password := generator.GeneratePassword()
	assert.NotEmpty(t, password)
	assert.Len(t, password, 10)
}

func TestGeneratorGeneratePasswordWithMinMaxLengthShouldSucceed(t *testing.T) {
	generatorConfig := generator.Config{
		AllowLowerCaseLetters: true,
		MinLength:             10,
		MaxLength:             20,
	}
	generator, _ := generator.New(&generatorConfig)

	for i := 0; i < 100; i++ {
		password := generator.GeneratePassword()
		assert.NotEmpty(t, password)
		assert.GreaterOrEqual(t, len(password), 10)
		assert.LessOrEqual(t, len(password), 20)
	}
}
