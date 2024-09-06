package functions

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Calcul(s string) {
	var data []float64
	tab_data := strings.Split(s, "\n")
	for _, r := range tab_data {
		if r == "" {
			continue
		}
		temp, err := strconv.Atoi(r)
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, float64(temp))
		temp = 0
	}
	sort.Float64s(data)

	fmt.Println("Average:", math.Round(Average(data)))
	fmt.Println("Median:", math.Round(Median(data)))
	fmt.Printf("Variance: %.0f \n", math.Round(Variance(data, Average(data))))
	fmt.Println("Standard Deviation:", math.Round(Standard_Deviation(Variance(data, Average(data)))))
}
