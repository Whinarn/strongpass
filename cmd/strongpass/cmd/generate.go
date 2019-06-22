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
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,g
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/whinarn/strongpass/pkg/generator"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a strong password",
	Long:  "Generates a strong password with your requirements.",
	Run: func(cmd *cobra.Command, args []string) {
		config := getGeneratorConfig()
		generator, err := generator.New(config)
		if err != nil {
			log.Fatal(err)
			return
		}

		password := generator.GeneratePassword()
		fmt.Println(password)
	},
}
var generateCharSet string
var generateLowerCaseLetters bool
var generateUpperCaseLetters bool
var generateDigits bool
var generateSpecials bool
var generateLength int
var generateMinLength int
var generateMaxLength int
var generateMinLowerCaseLetters int
var generateMinUpperCaseLetters int
var generateMinDigits int
var generateMinSpecials int
var generateMinShuffleCount int
var generateMaxShuffleCount int

func init() {
	generateCmd.Flags().StringVarP(&generateCharSet, "charset", "c", "", "The custom charset to use")
	generateCmd.Flags().BoolVarP(&generateLowerCaseLetters, "lowercase", "l", true, "The generator will used lower-case letters")
	generateCmd.Flags().BoolVarP(&generateUpperCaseLetters, "uppercase", "u", true, "The generator will used upper-case letters")
	generateCmd.Flags().BoolVarP(&generateDigits, "digits", "d", true, "The generator will use digits")
	generateCmd.Flags().BoolVarP(&generateSpecials, "specials", "s", true, "The generator will use special symbols")
	generateCmd.Flags().IntVar(&generateLength, "len", 0, "The length of the password, overrides minimum and maximum")
	generateCmd.Flags().IntVar(&generateMinLength, "min", 20, "The minimum length of the password")
	generateCmd.Flags().IntVar(&generateMaxLength, "max", 26, "The maximum length of the password")
	generateCmd.Flags().IntVar(&generateMinLowerCaseLetters, "minlowercase", 1, "The minumum number of lower-case letters in the password")
	generateCmd.Flags().IntVar(&generateMinUpperCaseLetters, "minuppercase", 1, "The minumum number of upper-case letters in the password")
	generateCmd.Flags().IntVar(&generateMinDigits, "mindigits", 1, "The minumum number of digits in the password")
	generateCmd.Flags().IntVar(&generateMinSpecials, "minspecials", 1, "The minumum number of special symbols in the password")
	generateCmd.Flags().IntVar(&generateMinShuffleCount, "minshuffle", 4, "The minumum number of random shuffles")
	generateCmd.Flags().IntVar(&generateMaxShuffleCount, "maxshuffle", 10, "The maximum number of random shuffles")
	rootCmd.AddCommand(generateCmd)
}

func getGeneratorConfig() *generator.Config {
	if generateLength > 0 {
		generateMinLength = generateLength
		generateMaxLength = generateLength
	}

	return &generator.Config{
		CharSet:               []rune(generateCharSet),
		AllowLowerCaseLetters: generateLowerCaseLetters,
		AllowUpperCaseLetters: generateUpperCaseLetters,
		AllowDigits:           generateDigits,
		AllowSpecials:         generateSpecials,
		MinLength:             generateMinLength,
		MaxLength:             generateMaxLength,
		MinLowerCaseLetters:   generateMinLowerCaseLetters,
		MinUpperCaseLetters:   generateMinUpperCaseLetters,
		MinDigits:             generateMinDigits,
		MinSpecials:           generateMinSpecials,
	}
}
