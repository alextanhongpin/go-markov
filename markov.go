package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
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
		c.chain[p.String()][string(b)]++
		p.Shift(b)
	}
}

func (c *Chain) Generate(w io.Writer, n int) {
	bw := bufio.NewWriter(w)

	var key string
	for key = range c.chain {
		break
	}
	p := NewPrefix(2)
	var total int
	for i := 2; i < n; i++ {
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
		bw.WriteString(" ")
		p.Shift([]byte(key))
	}

	bw.Flush()
}

func main() {
	c := NewChain()
	c.Build(os.Stdin)
	// fmt.Println(c.chain)
	c.Generate(os.Stdout, 200)
}
