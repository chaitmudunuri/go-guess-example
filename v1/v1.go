package guessing

import (
	"fmt"
	"math/rand"
	"time"
)

func Exec() {
	fmt.Println("Guess the number")
	guess := 0
	secret := rand.New(rand.NewSource(time.Now().Unix())).Intn(100)

	for {
		if _, err := fmt.Scan(&guess); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Input: ", guess, secret)

		if guess > secret {
			fmt.Println("Try smaller!")
		} else if guess < secret {
			fmt.Println("Try bigger!")
		} else {
			fmt.Println("Yay!")
			break
		}
	}
}
