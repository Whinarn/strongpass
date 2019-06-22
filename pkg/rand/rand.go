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

// Credits: https://web.archive.org/web/20190622035629/https://yourbasic.org/golang/crypto-rand-int/

package rand

import (
	crand "crypto/rand"
	"encoding/binary"
	"log"
	rand "math/rand"
)

type cryptoSource struct{}

var globalRand = New()

// New returns a new cryptographically secure pseudo-random number generator.
func New() *rand.Rand {
	return rand.New(&cryptoSource{})
}

// Int returns a non-negative pseudo-random int.
func Int() int {
	return globalRand.Int()
}

// Intn returns, as an int, a non-negative pseudo-random number in [0,n).
// It panics if n <= 0.
func Intn(n int) int {
	return globalRand.Intn(n)
}

// Shuffle pseudo-randomizes the order of elements.
// n is the number of elements. Shuffle panics if n < 0.
// swap swaps the elements with indexes i and j.
func Shuffle(n int, swap func(i, j int)) {
	globalRand.Shuffle(n, swap)
}

func (s *cryptoSource) Seed(seed int64) {}

func (s *cryptoSource) Int63() int64 {
	return int64(s.Uint64() &^ uint64(1<<63))
}

func (s *cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
