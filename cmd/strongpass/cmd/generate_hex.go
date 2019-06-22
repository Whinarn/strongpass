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
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var generateHexCmd = &cobra.Command{
	Use:   "generate-hex",
	Short: "Generates a strong password",
	Long:  "Generates a strong password with your requirements.",
	Run: func(cmd *cobra.Command, args []string) {
		if generateHexLength <= 0 {
			log.Fatal("The length must be over zero.")
			return
		}

		buffer := make([]byte, generateHexLength)
		_, err := rand.Read(buffer)
		if err != nil {
			log.Fatal(errors.Wrap(err, "Failed to generate random bytes"))
			return
		}

		password := hex.EncodeToString(buffer)
		if generateHexUpper {
			password = strings.ToUpper(password)
		}

		fmt.Println(password)
	},
}
var generateHexUpper bool
var generateHexLength int

func init() {
	generateHexCmd.Flags().BoolVarP(&generateHexUpper, "uppercase", "u", false, "The generator will use upper-cased hexadecimals")
	generateHexCmd.Flags().IntVarP(&generateHexLength, "len", "l", 32, "The length of the password in bytes")
	rootCmd.AddCommand(generateHexCmd)
}
