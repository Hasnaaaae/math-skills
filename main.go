package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("\x1b[33mUsage: go run . data.txt\033[0m")
		return
	}
	dataR, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("the file is ")
		log.Fatal(0)
	}
	Calc(string(dataR))
}

func Calc(s string) {
	var tab_int []int
	tab_data := strings.Split(s, "\n")
	for i := 0; i < len(tab_data); i++ {
		temp, err := strconv.Atoi(tab_data[i])
		if err != nil {
			log.Fatal(1)
		}
		tab_int = append(tab_int, temp)
		temp = 0
	}
	fmt.Println("Average:", Average(tab_int))
	fmt.Println("Median:", Median(tab_int))
	fmt.Println("Variance:", Variance(tab_int, Average(tab_int)))
	fmt.Println("Standard Deviation:", Standard_Deviation(Variance(tab_int, Average(tab_int))))
}

func Average(tab []int) int {
	var average, somme int
	for i := 0; i < len(tab); i++ {
		somme = somme + tab[i]
	}
	average = somme / len(tab)
	return average
}

func Median(tab []int) int {
	var mediam int
	if len(tab)%2 == 0 {
		mediam = (tab[(len(tab)/2)-1] + tab[len(tab)/2]) / 2
	} else {
		mediam = tab[len(tab)/2]
	}
	return mediam
}

func Variance(tab []int, average int) int {
	var variance, temp int
	var temp2 float64
	for i := 0; i < len(tab); i++ {
		temp2 = math.Pow(float64(tab[i]-average), 2)
		temp = int(temp2)
		variance = variance + temp
	}
	variance = variance / len(tab)
	return variance
}

func Standard_Deviation(variance int) int {
	var sd int
	temp := 0.0
	temp = math.Sqrt(float64(variance))
	sd = int(temp)
	return sd
}
