package main

import "fmt"

func main() {

	menu := map[string]float64{
		"pizza" : 40.00,
		"suco" : 12.00,
		"x-tudo": 22.90,
	}

	fmt.Println(menu)

	for key, value := range menu {
		fmt.Println(key, "R$", value)
	}
	contanova := novaConta("Astrubal")
	fmt.Println(contanova)

}