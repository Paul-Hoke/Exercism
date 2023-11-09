package chance

import "math/rand"
import "time"

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(20 - 1 + 1)
	//panic("Please implement the RollADie function")
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	rand.Seed(time.Now().UnixNano())
	return 0 + rand.Float64() * 12
	//panic("Please implement the GenerateWandEnergy function")
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
	return animals
	//panic("Please implement the ShuffleAnimals function")
}
