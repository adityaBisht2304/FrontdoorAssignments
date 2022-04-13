package course1week3

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Question_1() {
	arr := make([]int, 0)
	i := 0
	for true {
		str := ""
		fmt.Printf("Enter integer number or 'X' to exit\n")
		fmt.Scan(&str)

		if strings.ToLower(str) == "x" {
			break
		}
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Please enter number correctly\n")
		} else {
			arr = append(arr, num)
			i++
			sort.Ints(arr)
			printSlice(arr)
		}
	}
}

func printSlice(arr []int) {

	fmt.Println("Printing Sorted Slice")
	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}
