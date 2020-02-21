package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Dice Game 1
// 
// Write a dice rolling game where the goal is to roll dice until
// the sum of the dice is no more than 21.
//
// Make the game keep rolling a 6 sided die while the sum is less than
// 21. If the score is 21, then the player wins. If the score goes over
// 21, then the player loses. Let the player know the outcome of each
// dice roll and the outcome of the game.
// 
// Example:
// The player's score starts at 0. The die is rolled and the result is
// 3. The players score is now 3. Since the score is less than 3, the
// die is rolled again. The next result is 4. The player's score is now
// 7. The die keeps getting rolled and the next two results are 5 and 6.
// The player's score is 18.
// On the next roll, the score is 3 and the player's score is 21. The
// game prints "You win!" and quits.
// If however, the the roll was a 4, the player's score would have been
// 22 and the game would have printed "You lose!", then quit.

func main() {
	fmt.Println("You rolled:", roll())
}

// ================= You don't need to change anything below this line. ================


func init() {
	rand.Seed(time.Now().UnixNano())
}

func roll() int {
	return rand.Intn(6) + 1
}
