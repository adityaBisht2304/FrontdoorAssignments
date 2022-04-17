package course2week1

import "fmt"

func Question_1() {
	var nums []int = []int{}
	fmt.Printf("Enter 10 integers\n")
	for i := 0; i < 10; i++ {
		num := 0
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	BubbleSort(nums)
	Print(nums)
}

func Print(arr []int) {
	fmt.Printf("Printing Sorted array\n")
	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
}

func BubbleSort(arr []int) {
	length := len(arr)

	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			swap(arr, j)
		}
	}
}

func swap(arr []int, index int) {
	if arr[index] > arr[index+1] {
		temp := arr[index]
		arr[index] = arr[index+1]
		arr[index+1] = temp
	}
}
