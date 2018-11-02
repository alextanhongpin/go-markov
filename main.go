package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

var separator = []byte(" ")

type Prefix struct {
	prev, next []byte
}

func (p Prefix) String() string {
	return string(append(p.prev, append(separator, p.next...)...))
}

func (p *Prefix) Shift(b []byte) {
	p.prev = p.next
	p.next = b
}

type Chain struct {
	chain  map[string]map[string]int
	prefix Prefix
}

func NewChain() *Chain {
	return &Chain{
		chain: make(map[string]map[string]int),
		prefix: Prefix{
			prev: []byte(""),
			next: []byte(""),
		},
	}
}

func (c *Chain) Build(r io.Reader) {
	br := bufio.NewReader(r)
	for {
		var b []byte
		if _, err := fmt.Fscan(br, &b); err != nil {
			break
		}

		key := c.prefix.String()
		if _, exist := c.chain[key]; exist {
			c.chain[key][string(b)]++
		} else {
			c.chain[key] = map[string]int{
				string(b): 1,
			}
		}
		c.prefix.Shift(b)
	}
}

func (c *Chain) Generate(n int) string {
	p := Prefix{
		prev: []byte(""),
		next: []byte(""),
	}
	var result []string
	var total int
	var key string
	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		total = 0
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
		result = append(result, key)
		p.Shift([]byte(key))
	}
	return strings.Join(result, " ")
}

func main() {
	c := NewChain()

	str := `Never was a leader
Never had a thing for fairytales
Not really a believer, oh-oh
Small voice in the quiet
Guess I never dared to know myself
Can my heart beat quiet? No
But then there was you (but then there was you)
Yeah, then there was you
Pull me out of the crowd
You were telling the truth (you were telling the truth)
Yeah (yeah, yeah)
I got something to say now
'Cause you tell me that there's no way I couldn't go
Nothing I couldn't do
Yeah
I want to get louder
I got to get louder
We 'bout to go up baby, up we go
We 'bout to go up baby, up we go
We're blowing out speakers
Our heart a little clearer
We 'bout to go up baby, up we go
We 'bout to go up baby, up we go
For worse or for better
Gonna give it to you
In capital letters
We put a crack in the shadows
And you tell me it's okay to be the light
And not to swim in the shallows
No, no
And I wanna get drunk with you
When we lie so still, but you're taking me places
Holding me onto you
And we don't care who's watching us, baby
But then there was you
(But then there was you)
Yeah, then there was you
Pull me out of the crowd
You were telling the truth
(You were telling the truth)
Yeah (yeah, yeah)
I got something to say now
'Cause you tell me that there's no way I couldn't go
Nothing I couldn't do (no, no, no, no)
Yeah
I want to get louder
I got to get louder
We 'bout to go up baby, up we go
We 'bout to go up baby, up we go
We're blowing out speakers
Our heart a little clearer
We 'bout to go up baby, up we go
We 'bout to go up baby, up we go
For worse or for better
Gonna give it to you
In capital letters
In capital letters
In capital–
Gonna give it to you
Gonna give it to you
Gonna give it to you
I want to get louder
I got to get louder
We 'bout to go up baby, up we go
We 'bout to go up baby, up we go
We're blowing out speakers
Our heart a little clearer
We 'bout to go up baby, up we go
We 'bout to go up baby, up we go
I want to get louder
I got to get louder
We 'bout to go up baby, up we go
We 'bout to go up baby, up we go
We're blowing out speakers
Our heart a little clearer
We 'bout to go up baby, up we go
We 'bout to go up baby, up we go
For worse or for better
Gonna give it to you
In capital letters

You made plans and I, I made problems
We were sleeping back to back
We know this thing wasn't built to last and
Good on paper, picture perfect
Chased the high too far, too fast
Picket white fence, but we paint it black
Ooh, and I wished you had hurt me harder than I hurt you
Ooh, and I wish you wouldn't wait for me but you always do
I've been hoping somebody loves you in the ways I couldn't
Somebody's taking care of all of the mess I've made
Someone you don't have to change
I've been hoping
Someone will love you, let me go
Someone will love you, let me go
I've been hoping
Someone will love you, let me go
It's been some time, but this time ain't even
I can leave it in the past
But you're holding on to what you never had
Good on paper, picture perfect
Chased the high too far, too fast
Picket white fence, but we paint it black
Ooh, and I wished you had hurt me harder than I hurt you
Ooh, and I wish you would have waited for me but you always do
I've been hoping somebody loves you in the ways I couldn't
Somebody's taking care of all of the mess I've made
Someone you don't have to change
I've been hoping
Someone will love you, let me go
Someone will love you, let me go
I've been hoping
Someone will love you, let me go (go, go, go)
Someone will love you, let me go (go, go, go)
Someone will love you, let me go (go, go, go)
Someone will love you, let me go (go, go, go)
Someone will love you, let me go
I've been hoping somebody loves you in the ways I couldn't
Somebody's taking care of all of the mess I've made
Someone you don't have to change
I've been hoping
Someone will love you, let me go

You know just what to say
Shit, that scares me, I should just walk away
But I can't move my feet
The more that I know you, the more I want to
Something inside me's changed
I was so much younger yesterday, oh
I didn't know that I was starving till I tasted you
Don't need no butterflies when you give me the whole damn zoo
By the way, by the way, you do things to my body
I didn't know that I was starving till I tasted you
By the way, by the way, you do things to my body
I didn't know that I was starving till I tasted you
You know just how to make my heart beat faster
Emotional earthquake, bring on disaster
You hit me head-on, got me weak in my knees
Yeah, something inside me's changed
I was so much younger yesterday, ye-eah
So much younger yesterday, oh, yeah
I didn't know that I was starving till I tasted you
Don't need no butterflies when you give me the whole damn zoo
By the way, by the way, you do things to my body
I didn't know that I was starving till I tasted you
By the way, by the way, you do things to my body
I didn't know that I was starving till I tasted you
You, yeah, till I tasted you
(I didn't know that I-I didn't know that I-till I tasted you)
By the way, by the way, you do things to my body
I didn't know that I was starving till I tasted you, ooh, ooh, ooh, ooh
Na-na-na-na
Na-na-na-na
The more that I know you, the more I want to
Something inside me's changed
I was so much younger yesterday
`

	str = strings.ToLower(str)
	c.Build(strings.NewReader(str))
	fmt.Println(c.chain)
	fmt.Println(c.Generate(100))
}
