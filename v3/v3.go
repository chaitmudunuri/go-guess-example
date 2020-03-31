package guessing

import (
	"fmt"
	"math/rand"
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
	return fmt.Sprintf("%dt [%d - %d]", g.tries, g.left, g.right)
}

func Exec() {
	g := Guess{guess: 0, tries: 0, min: 0, max: 100, left: 0, right: 100,
		secret: rand.New(rand.NewSource(time.Now().Unix())).Intn(100)}

	for !g.read().ok() {
		fmt.Println(g)
	}

	fmt.Printf("Yay! Got it in %d tries.\n", g.tries)
}
