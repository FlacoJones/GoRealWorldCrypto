package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
)

func bday() {
	n := 23
	maxHashValue := big.NewInt(365)
	collisionCounter := 0
	generations := 1000

	var wg sync.WaitGroup
	collisionChan := make(chan bool, generations)

	for j := 0; j < generations; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			seenHashes := make(map[string]bool)
			for i := 0; i < n; i++ {
				randomHash, _ := rand.Int(rand.Reader, maxHashValue)
				if _, exists := seenHashes[randomHash.String()]; exists {
					collisionChan <- true
					return
				}
				seenHashes[randomHash.String()] = true
			}
			collisionChan <- false
		}()
	}

	go func() {
		wg.Wait()
		close(collisionChan)
	}()

	for collision := range collisionChan {
		if collision {
			collisionCounter++
		}
	}

	collisionPercentage := float64(collisionCounter) / float64(generations) * 100
	fmt.Printf("Out of %d random groups of 23 students with 365 random birthdays, collisions were found %.2f%% of the time", generations, collisionPercentage)
	fmt.Println()
}
