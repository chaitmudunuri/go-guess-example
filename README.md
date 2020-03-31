# Go Implementation of Rust Guessing Game 
## Learn Rust Basics using the game

We start with the basic code from the **Rust Book** and add new data structures and functions to make the code modular and reusable.

### Version 1  
This is straight replication from Rust version 1 implementation.  
User is asked to guess a number between 1 and 100.  
If the guess is bigger than secret, it says _Too big!_ and it small says _Too small!_

**Source**: [Rust Guess Example](https://github.com/chaitmudunuri/rust-guess-example)  

### Version 2  
Version 2 adds more context by adding number of tries and range information using min and max values.    

Added variables `tries`, `min` and `max`

```
guess := 0
tries := 0
min := 0
max := 100
```

**Sample Output**
```
Guess the number
50
Try smaller! [0 - 50] Guessed: 50 in 1 tries.
25
Try bigger! [25 - 50] Guessed: 25 in 2 tries.
37
Try smaller! [25 - 37] Guessed: 37 in 3 tries.
30
Try smaller! [25 - 30] Guessed: 30 in 4 tries.
28
Yay! Got it in 5 tries.
```

### Version 3  
Version 3 adds `Guess` struct to keep all the guess related information at one place.    
* Implements `read()` and `ok()` functions on `Guess`
* Implements `String()` to print the struct using `fmt::Print` functions

```
type Guess struct {
	guess, tries, secret, min, max, left, right int
}
```

**Sample Output**
```
Guess? 50
1t [0 - 50]
Guess? 25
2t [25 - 50]
Guess? 37
Yay! Got it in 3 tries.
```

### Version 4  
Verson 4 implements guessing range indication like Rust final version.

```
fn main() {
    let mut guessing_state = GuessingState::new(1, 100);
    while !guessing_state.input().ok() {}
}
```

* `GuessingState` will implement `new` and `Display` trait
* `new` will take guessing range as input so that it is reusable
* `Display` is like `.toString` in Java
* Display will show a guessing range instead of Debug input
* Display will be like `[min]========[left]====[right]========[max]`
* `read_input()` and `check_guess()` will be implemented on `GuessingState`

**Sample Output**
```
Guess? 50
[=================================================]==================================================
Guess? 37
=====================================[============]==================================================
Guess? 42
==========================================[=======]==================================================
Guess? 48
=============================================[==]====================================================
Guess? 47
Yay! Got it in 4 tries.
```
