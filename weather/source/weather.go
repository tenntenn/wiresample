package source

// Celsius represents temperature by the celsius.
type Celsius float64

// Weather represents weather infomation.
type Weather struct {
	Temperature       Celsius
	PrecipProbability float64
	Summary           string
}
