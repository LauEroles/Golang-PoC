package main

import "fmt"

func main() {

	var hello string = "Hello variable World"

	// los : sirven para declarar la variable y asignar el valor al mismo tiempo
	bye := "Bye bye"

	numbers := []int{
		1, 2, 5, 30, 19, 55, 2, 4, 3, 32, 28, 78, 59, 11, 18}

	fmt.Println(hello)

	showItems(numbers)

	sortedNumbers := bubbleSort(numbers)

	showItems(sortedNumbers)
	showItems(numbers)

	fmt.Println(bye)

	num := int16(1)

	fmt.Println(num)

}

//  los arrays que no son iniciailizados con numero determinado se llaman slices en Go
// vamos a hacer una funcion que muestre los items de un array
func showItems(items []int) {

	for i := 0; i < len(items); i++ {
		item := items[i]
		fmt.Printf("%d, ", item)
	}
	fmt.Println("")

}

func bubbleSort(unsortedItems []int) []int {
	// Copy the original array into a new one. This allows to preserve the unsorted values
	items := make([]int, len(unsortedItems))
	copy(items, unsortedItems)

	for i := 0; i < len(items); i++ {
		swapped := false
		for j := len(items) - 1; j > i; j-- {
			if items[j] < items[j-1] {

				// Traditional swapping method with an auxiliary variable
				/*aux := items[j]
				items[j] = items[j-1]
				items[j-1] = aux*/

				// The golang SUPER swapping method, no additional variable needed (WOWWWW)
				items[j], items[j-1] = items[j-1], items[j]

				swapped = true
			}

		}
		if !swapped {
			break
		}
	}

	return items
}
