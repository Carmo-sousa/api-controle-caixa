package main

import "fmt"

func main() {
	foods := []string{"Banana", "Maçã", "Mamão"}
	pos := 0
	fmt.Println(foods)
	foods = remove(foods, pos)
	fmt.Println(foods)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
