package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	guess := Guess{guess: 0, tries: 0, min: 0, max: 100, left: 0, right: 100,
		secret: rand.New(rand.NewSource(time.Now().Unix())).Intn(100)}
	fmt.Println(guess)

	mux := http.NewServeMux()
	mux.Handle("/guess", http.HandlerFunc(guess.check))
	log.Fatal(http.ListenAndServe(":7000", mux))
}

type Guess struct {
	guess, tries, secret, min, max, left, right int
}

// Sample Request - https://localhost:7000/guess?guess=90
func (g *Guess) check(w http.ResponseWriter, req *http.Request) {
	input, err := strconv.Atoi(req.URL.Query().Get(("guess")))

	if err != nil {
		WriteJSON(w, http.StatusForbidden, err.Error())
	} else if input == g.secret {
		WriteJSON(w, http.StatusOK, "{result: Yay!}")
	} else if input < g.secret {
		WriteJSON(w, http.StatusOK, "{result: Try bigger!}")
	} else if input > g.secret {
		WriteJSON(w, http.StatusOK, "{result: Try smaller!}")
	} else {
		WriteJSON(w, http.StatusForbidden, "{result: Something wrong!}")
	}
}

func WriteJSON(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, msg)
}
