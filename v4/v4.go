package guessing

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Guess struct {
	guess, tries, secret, min, max, left, right int
}

func (g *Guess) read() *Guess {
	fmt.Printf("Guess? ")
	if _, err := fmt.Scan(&g.guess); err != nil {
		fmt.Println(err)
	}

	return g
}

func (g *Guess) ok() bool {
	g.tries++
	if g.guess > g.secret {
		g.right = g.guess
		return false
	} else if g.guess < g.secret {
		g.left = g.guess
		return false
	} else {
		return true
	}
}

func (g Guess) String() string {
	var buffer strings.Builder
	for i := g.min; i <= g.max; i++ {
		if i == g.left {
			buffer.WriteByte('[')
		} else if i == g.right {
			buffer.WriteByte(']')
		} else {
			buffer.WriteByte('=')
		}
	}

	return fmt.Sprintf(buffer.String())
}

func Exec() {
	guess := Guess{guess: 0, tries: 0, min: 0, max: 100, left: 0, right: 100,
		secret: rand.New(rand.NewSource(time.Now().Unix())).Intn(100)}

	for !guess.read().ok() {
		fmt.Println(guess)
	}

	fmt.Printf("Yay! Got it in %d tries.\n", guess.tries)
}
