package main

import (
	"crypto/rand"
	"math/big"
	"sync"
	"testing"
)

func birthdayCollision(generations int, n int, maxHashValue *big.Int, multiThreaded bool) float64 {
	collisionCounter := 0

	if multiThreaded {
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
	} else {
		for j := 0; j < generations; j++ {
			seenHashes := make(map[string]bool)
			for i := 0; i < n; i++ {
				randomHash, _ := rand.Int(rand.Reader, maxHashValue)
				if _, exists := seenHashes[randomHash.String()]; exists {
					collisionCounter++
					break
				}
				seenHashes[randomHash.String()] = true
			}
		}
	}

	return float64(collisionCounter) / float64(generations) * 100
}

func BenchmarkMultiThreaded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		birthdayCollision(100, 23, big.NewInt(365), true)
	}
}

func BenchmarkSingleThreaded(b *testing.B) {
	for i := 0; i < b.N; i++ {
		birthdayCollision(100, 23, big.NewInt(365), false)
	}
}
