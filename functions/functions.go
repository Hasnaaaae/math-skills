package functions

import "math"

// Average function

func Average(tab []float64) float64 {
	var average, somme float64
	for i := 0; i < len(tab); i++ {
		somme = somme + tab[i]
	}
	average = somme / float64(len(tab))
	return average
}

// Median function

func Median(tab []float64) float64 {
	var mediam float64
	if len(tab)%2 == 0 {
		mediam = (tab[(len(tab)/2)-1] + tab[len(tab)/2]) / 2
	} else {
		mediam = tab[len(tab)/2]
	}
	return mediam
}

// Variance function

func Variance(tab []float64, average float64) float64 {
	var variance float64
	for i := 0; i < len(tab); i++ {
		variance += math.Pow(tab[i]-average, 2)
	}
	variance = variance / float64(len(tab))
	return variance
}

// Standard_Deviation function

func Standard_Deviation(variance float64) float64 {
	sd := math.Sqrt(variance)
	return sd
}
