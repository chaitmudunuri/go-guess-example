package guessing

import (
	"fmt"
	"math/rand"
	"time"
)

func Exec() {
	fmt.Println("Guess the number")
	guess := 0
	tries := 0
	min := 0
	max := 100
	secret := rand.New(rand.NewSource(time.Now().Unix())).Intn(100)

	for {
		if _, err := fmt.Scan(&guess); err != nil {
			fmt.Println(err)
		}
		tries++

		if guess > secret {
			fmt.Printf("Try smaller! ")
			max = guess
		} else if guess < secret {
			fmt.Printf("Try bigger! ")
			min = guess
		} else {
			fmt.Printf("Yay! Got it in %d tries.\n", tries)
			break
		}
		fmt.Printf("[%d - %d] Guessed: %d in %d tries.\n", min, max, guess, tries)
	}
}
