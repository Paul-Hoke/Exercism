package resistorcolortrio

import (
	"strconv"
	"fmt"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	
	colorMap := PopulateColorMap()
	
	rawString := CalculateRawString(colors, colorMap)
	
	return rawString
	//panic("Implement the Label function")
}

func PopulateColorMap() map[string]int {
	return map[string]int {
		"black": 0,
		"brown": 1,
		"red": 2,
		"orange": 3,
		"yellow": 4,
		"green": 5,
		"blue": 6,
		"violet": 7,
		"grey": 8,
		"white": 9,
	}
}

func CalculateRawString(colors []string, inMap map[string]int) string {
	result := ""
	
	outerLabel:
	for i, v := range colors {
		
		switch {
		case i > 2:
			break outerLabel
		case i < 2:
			result = fmt.Sprintf("%s%s", result, strconv.Itoa(inMap[v]))
		default:
			result = fmt.Sprintf("%s%s", result, GetZeroes(inMap[v]))
		}
	}
	return result
}

func GetZeroes(colorValue int) string {
	result := ""
	
	for colorValue > 0 {
		result = fmt.Sprintf("%s%s", result, "0")
		colorValue--
	}
	return result
}
