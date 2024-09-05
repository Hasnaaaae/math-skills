package main

import (
	"fmt"
	"log"
	"os"

	"math_skills/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("\x1b[33mUsage: go run . data.txt\033[0m")
		return
	}
	dataR, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("the file is ")
		log.Fatal()
	}
	functions.Calcul(string(dataR))
}
