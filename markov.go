package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

// https://rosettacode.org/wiki/Markov_chain_text_generator#Go
// https://golang.org/doc/codewalk/markov/
type Prefix [][]byte

func NewPrefix(n int) Prefix {
	return make([][]byte, 2)
}

func (p Prefix) String() string {
	return string(bytes.Join(p, []byte(" ")))
}

func (p Prefix) Shift(b []byte) {
	copy(p, p[1:])
	p[len(p)-1] = b
}

func (p Prefix) ShiftString(s string) {
	str := strings.Split(s, " ")
	for _, s := range str {
		p.Shift([]byte(s))
	}
}

type Chain struct {
	chain       map[string]map[string]int
	capitalized int
}

func NewChain() *Chain {
	return &Chain{
		chain: make(map[string]map[string]int),
	}
}

func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)

	p := NewPrefix(2)

	for {
		var b []byte
		if _, err := fmt.Fscan(br, &b); err != nil {
			return
		}
		_, exist := c.chain[p.String()]
		if !exist {
			c.chain[p.String()] = make(map[string]int)
		}
		if isUpper(p.String()) {
			c.capitalized++
		}
		// Increment if it is capitalized. We can use this to get the
		// probability of the words that is capitalized.
		c.chain[p.String()][string(b)]++
		p.Shift(b)
	}
}

func (c *Chain) Generate(w io.Writer, n int) {
	bw := bufio.NewWriter(w)
	count := n
	// Take the first key. Ensure that it is capitalized.
	randomPrefix := rand.Intn(c.capitalized)
	var key string
	for key = range c.chain {
		if !isUpper(key) {
			continue
		}
		randomPrefix -= 2
		if randomPrefix < 0 {
			break
		}
	}
	p := NewPrefix(2)
	p.ShiftString(key)
	bw.WriteString(key)
	count -= 2

	var total int
	var hasPunct bool
	for {
		if count < 1 && hasPunct {
			break
		}
		bw.WriteString(" ")
		total = 0
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			return
		}
		for _, freq := range choices {
			total += freq
		}
		choice := rand.Intn(total)
		for key = range choices {
			if choice == 0 {
				break
			}
			choice--
		}
		bw.WriteString(key)
		p.Shift([]byte(key))
		if isPunct(key) {
			hasPunct = true
		} else {
			hasPunct = false
		}
		count -= 2
	}

	bw.Flush()
}

func isUpper(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}

func isPunct(s string) bool {
	r, _ := utf8.DecodeLastRuneInString(s)
	// return unicode.IsPunct(r)
	return r == '.' || r == '?' || r == '!'
}

func main() {
	c := NewChain()
	c.Build(os.Stdin)
	// fmt.Println(c.chain)
	c.Generate(os.Stdout, 200)
}
