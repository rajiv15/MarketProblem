//Package main contains user interface logic for pruchasing the products from the market.
package main

import (
	"bufio"
	basket "checkoutsystem/basket"
	co "checkoutsystem/checkout"
	product "checkoutsystem/product"
	"fmt"
	"os"
	"strings"
)

//clearBasket refresh basket to start including new items
func clearBasket() {
	for k := range *basket.Basket.GetBasket() {
		delete(*basket.Basket.GetBasket(), k)
	}
}

//showItems helper function for showing the available product in the market.
func showItems() {

	fmt.Println("\nFarmers Market Items")

	fmt.Printf("Product Code\tName   \tPrice\n")
	for _, item := range product.Products {
		fmt.Printf("%-12s\t%4s\t%f\n", item.ProductCode, item.Name, item.Price)
	}
}

//displayMenu helper funtion for showing the menu for purchasing the product.
func displayMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Println("The Farmer's Market")
		fmt.Println("-------------------")
		fmt.Println("1. Insert new item into the basket")
		fmt.Println("2. Show register")
		fmt.Println("3. Clear Basket")
		fmt.Println("4. Exit")

		fmt.Println()
		fmt.Println("Enter the Choice: ")
		choice := 0
		_, err := fmt.Scanf("%d", &choice)
		fmt.Println()

		if err != nil {
			fmt.Println("Error: Invalid Input. Expected Integer")
			reader.ReadString('\n')
			fmt.Println()
		}

		switch choice {
		case 1:
			showItems()
			fmt.Println("\nEnter the Product Code to add it to Basket: ")
			itemID, _ := reader.ReadString('\n')
			itemID = strings.ToUpper(strings.TrimRight(itemID, "\n"))

			if _, ok := product.Products[itemID]; !ok {
				fmt.Println("Error: Invalid Product Code. Enter again")
				fmt.Println()
				continue
			}
			basket.Basket.AddItem(strings.ToUpper(itemID))
		case 2:
			co.Cos.PrintBills()
		case 3:
			clearBasket()
		default:
			os.Exit(0)
		}
	}
}

func main() {
	displayMenu()
}
