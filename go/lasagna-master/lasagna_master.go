package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, perparationTime int) int {
	if(perparationTime == 0) {
		return len(layers) * 2
	}
	return len(layers) * perparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	
	for _, v := range layers {
		
		switch v {
		case "sauce":
			sauce += 0.2
		case "noodles":
			noodles += 50
		}
		
	}
	
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numPortions int) []float64 {
	result := make([]float64, 0)
	for _, v := range quantities {
		result = append(result, (v / 2.0) * float64(numPortions))
	}
	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
