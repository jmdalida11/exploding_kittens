package game

import "math/rand"

func randomizer[T any](items []T) {
	rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
}
