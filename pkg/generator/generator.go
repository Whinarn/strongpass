/*
MIT License

Copyright(c) 2019 Mattias Edlund

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package generator

import (
	"github.com/pkg/errors"
	"github.com/whinarn/strongpass/pkg/rand"
)

var (
	lowerCaseLetterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	upperCaseLetterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digitRunes           = []rune("0123456789")
	specialRunes         = []rune("§½!#¤%&/()[]{}=?+-*\\£$~^.,:;_<>|@")
)

// Generator is a password generator.
type Generator struct {
	charSet             []rune
	minLength           int
	maxLength           int
	minLowerCaseLetters int
	minUpperCaseLetters int
	minDigits           int
	minSpecials         int
	minShuffleCount     int
	maxShuffleCount     int
}

// Config is the password generator configuration.
type Config struct {
	CharSet []rune

	AllowLowerCaseLetters bool
	AllowUpperCaseLetters bool
	AllowDigits           bool
	AllowSpecials         bool

	MinLength           int
	MaxLength           int
	MinLowerCaseLetters int
	MinUpperCaseLetters int
	MinDigits           int
	MinSpecials         int

	MinShuffleCount int
	MaxShuffleCount int
}

// New returns a new generator. If config is nil, the default configuration is used.
func New(config *Config) (*Generator, error) {
	if config == nil {
		config = DefaultConfig()
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	charSet := config.prepareCharSet()
	return &Generator{
		charSet:             charSet,
		minLength:           config.MinLength,
		maxLength:           config.MaxLength,
		minLowerCaseLetters: config.MinLowerCaseLetters,
		minUpperCaseLetters: config.MinUpperCaseLetters,
		minDigits:           config.MinDigits,
		minSpecials:         config.MinSpecials,
		minShuffleCount:     config.MinShuffleCount,
		maxShuffleCount:     config.MaxShuffleCount,
	}, nil
}

// DefaultConfig returns the default configuration.
func DefaultConfig() *Config {
	return &Config{
		AllowLowerCaseLetters: true,
		AllowUpperCaseLetters: true,
		AllowDigits:           true,
		AllowSpecials:         true,
		MinLength:             20,
		MaxLength:             26,
		MinLowerCaseLetters:   1,
		MinUpperCaseLetters:   1,
		MinDigits:             1,
		MinSpecials:           1,
		MinShuffleCount:       4,
		MaxShuffleCount:       10,
	}
}

// GeneratePassword generates a password.
func (gen *Generator) GeneratePassword() string {
	length := gen.minLength
	if gen.maxLength > length {
		length += rand.Intn((gen.maxLength - length) + 1)
	}

	passwordChars := make([]rune, 0, length)
	passwordChars = gen.appendRandomChars(passwordChars, gen.minLowerCaseLetters, lowerCaseLetterRunes)
	passwordChars = gen.appendRandomChars(passwordChars, gen.minUpperCaseLetters, upperCaseLetterRunes)
	passwordChars = gen.appendRandomChars(passwordChars, gen.minDigits, digitRunes)
	passwordChars = gen.appendRandomChars(passwordChars, gen.minSpecials, specialRunes)

	remainingLength := length - len(passwordChars)
	passwordChars = gen.appendRandomChars(passwordChars, remainingLength, gen.charSet)

	shuffleCount := gen.minShuffleCount + rand.Intn((gen.maxShuffleCount-gen.minShuffleCount)+1)
	for i := 0; i < shuffleCount; i++ {
		rand.Shuffle(len(passwordChars), func(i, j int) {
			passwordChars[i], passwordChars[j] = passwordChars[j], passwordChars[i]
		})
	}

	return string(passwordChars)
}

func (gen *Generator) appendRandomChars(buffer []rune, length int, chars []rune) []rune {
	charsLen := len(chars)
	if charsLen == 0 {
		// There is nothing to append
		return buffer
	}

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(charsLen)
		buffer = append(buffer, chars[randomIndex])
	}
	return buffer
}

func (config *Config) validate() error {
	if config.MinLength <= 0 {
		return errors.New("The minumum length of a password cannot be zero or negative")
	}

	if config.MinLowerCaseLetters < 0 {
		config.MinLowerCaseLetters = 0
	}
	if config.MinUpperCaseLetters < 0 {
		config.MinUpperCaseLetters = 0
	}
	if config.MinDigits < 0 {
		config.MinDigits = 0
	}
	if config.MinSpecials < 0 {
		config.MinSpecials = 0
	}
	if config.MinShuffleCount < 1 {
		// There has to be at least 1 shuffle, otherwise there is no security at all
		config.MinShuffleCount = 1
	}
	if config.MaxShuffleCount < config.MinShuffleCount {
		config.MaxShuffleCount = config.MinShuffleCount
	}

	requiredMinimum := config.MinLowerCaseLetters + config.MinUpperCaseLetters +
		config.MinDigits + config.MinSpecials
	if config.MinLength < requiredMinimum {
		config.MinLength = requiredMinimum
	}
	if config.MaxLength < config.MinLength {
		config.MaxLength = config.MinLength
	}

	return nil
}

func (config *Config) prepareCharSet() []rune {
	var charSet []rune
	if len(config.CharSet) > 0 {
		// We have a pre-defined character set
		charSet = config.CharSet
	} else {
		// We have to create the character set
		if config.AllowLowerCaseLetters {
			charSet = append(charSet, lowerCaseLetterRunes...)
		}
		if config.AllowUpperCaseLetters {
			charSet = append(charSet, upperCaseLetterRunes...)
		}
		if config.AllowDigits {
			charSet = append(charSet, digitRunes...)
		}
		if config.AllowSpecials {
			charSet = append(charSet, specialRunes...)
		}
	}

	config.shuffleCharSet(charSet)
	return charSet
}

func (config *Config) shuffleCharSet(charSet []rune) {
	shuffleCount := config.MinShuffleCount + rand.Intn((config.MaxShuffleCount-config.MinShuffleCount)+1)
	for i := 0; i < shuffleCount; i++ {
		rand.Shuffle(len(charSet), func(i, j int) {
			charSet[i], charSet[j] = charSet[j], charSet[i]
		})
	}
}
