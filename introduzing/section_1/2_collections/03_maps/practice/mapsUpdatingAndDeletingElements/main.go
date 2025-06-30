package main

/*
	Create a map with product names and their prices.
	Update the price of one product, delete another,
	and print the final list of products and prices.
*/

import "fmt"

func main() {
	// Create a map with product names and their prices
	products := map[string]float64{
		"banana": 4.99,
		"apple":  3.50,
		"orange": 2.75,
	}

	// Update the price of banana
	products["banana"] = 5.49

	// Delete the orange from the map
	delete(products, "orange")

	// Print the updated products and their prices
	fmt.Println("Updated products:")
	for product, price := range products {
		fmt.Printf("%s: $%.2f\n", product, price)
	}
}
