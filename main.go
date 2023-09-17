// The Birthday Paradox and hash function

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Number of "people" (or in our case, hashes)
	n := 23 // The classic birthday problem uses 23 people

	maxHashValue := big.NewInt(365)

	// Upper bound for our "birthdays" (or in our case, hashes)
	// Imagine a small hash function that outputs values between 0 and 364
	// Approximately 2 ^ 8.511 possibilities, so there is a 50% chance that after 2 ^ (8.511 / 2) (between 16 and 32 random students) generations we'll collide
	// So we would expect that, half the time we check 23 hashes, we'll get a collision

	// Some may think 365/2

	// Rather than calculate all the possible combinations where people share a bday, find the odds that everyones bday is different

	// Odds of match plus odds of no match must add up to 100%
	// Match = 100 - no match
	// 364 / 365 * 363 / 365

	// 50.73% chance of one birthday match
	// Actually a large number of possible pair combination
	// Number of pairs grow quadratically

	// For demonstration, start with ten, then increase precision to show it's SLIGHTLY more likely to be ABOVE 50%

	// Bonus Point: Paralellize with Go

	collisionCounter := 0

	classrooms := 100

	for j := 0; j < classrooms; j++ {
		classroom := make(map[string]bool)

		for i := 0; i < n; i++ {
			// Generate a "birthday" (or in our case, a hash)
			studentBirthday, _ := rand.Int(rand.Reader, maxHashValue)

			// Check for collision
			if _, exists := classroom[studentBirthday.String()]; exists {
				fmt.Printf("Collision found after %d hashes\n", i+1)
				collisionCounter++
				break
			}

			// Store this hash in our set of seen hashes
			classroom[studentBirthday.String()] = true
		}
		fmt.Println("No collision found")
	}
	collisionPercentage := float64(collisionCounter) / float64(classrooms) * 100
	fmt.Printf("Out of %d random classrooms of 23 students with 365 random birthdays, collisions were found in %.2f%% of the classrooms", classrooms, collisionPercentage)
	fmt.Println()
}
