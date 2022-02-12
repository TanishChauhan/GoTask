package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var order1 string
	var order2 string
	var intVar int
	var arr = []int{}

	for true {
		fmt.Println("---------MENU----------")
		fmt.Print("C: coffee, 40rs\n D: dosa, 80 rs\n T: tomato soup, 20rs\n I : idli , 20rs\n V : vada, 25rs.\n B: bhature &chhole, 30rs\n P: paneer pakoda, 30rs\n M: manchurian, 90rs\n H: hakka noodle, 70rs.\n F: French fries, 30rs\n J: jalebi, 30rs\n L: lemonade, 15rs\n S: spring roll, 20rs\n")
		fmt.Println("-----------------------")
		fmt.Println("Press 'END' For closing the day.")
		fmt.Println("Please place your order :- ")
		fmt.Scan(&order1)
		order1 = strings.ToUpper(order1)

		if order1 != "END" {
			intVar, _ = strconv.Atoi(order1)

			fmt.Scan(&order2)
			order2 = strings.ToUpper(order2)

			bill := earned(intVar, order2)

			fmt.Println("++++++++++++++++++++")
			fmt.Println("Total bill: ", bill)
			fmt.Println("++++++++++++++++++++")
			arr = append(arr, bill)

		} else {
			DayEarning(arr)
		}

	}
}

func DayEarning(items []int) {
	var sum int = 0

	for i := 0; i < len(items); i++ {
		sum = items[i] + sum
	}
	fmt.Println("Total Income for the day is :  ", sum)
	os.Exit(0)
}

func earned(quantity int, order string) int {
	mon := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	mandy := quantity * int(mon[order])
	return int(mandy)
}
