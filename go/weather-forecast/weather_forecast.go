//Package weather can be used get current weather conditions at a given location
package weather

//CurrentCondition represents the current conditions at the location
var CurrentCondition string

//CurrentLocation represents the city we currently care about
var CurrentLocation string

//Forecast returns the current weather conditions given the city we are looking at
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
