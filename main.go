package main

import (
	"fmt"
)

func basicArray() {
	fmt.Println("\nBASIC ARRAY")
	var arr [2]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr)
}

func basicSlice() {
	fmt.Println("\nBASIC SLICE")
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"
	fmt.Println(cheeses)
}

func appendSlice() {
	fmt.Println("\nAPPEND SLICE ELEMENTS")
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"
	cheeses = append(cheeses, "Camembert", "Reblochon", "Picodon")
	fmt.Println(cheeses)
}

func deletingSlice() {
	fmt.Println("\nDELETE SLICE ELEMENT")
	var cheeses = make([]string, 3)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"
	cheeses[2] = "Camembert"
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
	// start at element and then delete until element
	cheeses = append(cheeses[:1], cheeses[2+1:]...)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
}

func copySliceElements() {
	fmt.Println("\nCOPY SLICE ELEMENTS")
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"
	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheeses)
	fmt.Println(smellyCheeses)
}

func copySingleSliceElements() {
	fmt.Println("\nCOPY SINGLE SLICE ELEMENTS")
	var cheeses = make([]string, 2)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"
	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheeses[1:])
	fmt.Println(smellyCheeses)
}

func maps() {
	fmt.Println("\nMAPS")
	// create a map with key of type string and value of type int
	var players = make(map[string]int)
	players["cook"] = 32
	players["bairstow"] = 27
	players["stokes"] = 26
	fmt.Println(players["cook"])
	fmt.Println(players["bairstow"])
}

func mapDelete() {
	fmt.Println("\nMAP DELETE")
	var players = make(map[string]int)
	players["cook"] = 32
	players["bairstow"] = 27
	delete(players, "cook")
	fmt.Println(players)
}

func main() {
	basicArray()
	basicSlice()
	appendSlice()
	deletingSlice()
	copySliceElements()
	copySingleSliceElements()
	maps()
	mapDelete()
}
