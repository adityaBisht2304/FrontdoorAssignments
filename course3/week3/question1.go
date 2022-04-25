package course3week3

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func Question_1() {
	fmt.Println("Enter space separated integers in a line. Add more than 4 integers")

	reader := bufio.NewReader(os.Stdin)
	byteArr, _, _ := reader.ReadLine()
	stringArr := strings.Split(string(byteArr), " ")
	arr := [4][]int{}
	length := len(stringArr)
	if length < 4 {
		fmt.Println("Array length is less than 4")
		return
	}
	//stringArr := []string{"3", "1", "2", "5", "4", "7"}
	i := 0
	for _, val := range stringArr {
		value, _ := strconv.Atoi(val)
		arr[i%4] = append(arr[i%4], value)
		i++
	}

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go SortArray(arr[i], i)
	}

	wg.Wait()

	tempArr1 := MergeSort(arr[0], arr[1])
	tempArr2 := MergeSort(arr[2], arr[3])

	tempArr := MergeSort(tempArr1, tempArr2)
	fmt.Printf("Printing complete sorted array\n")
	for _, val := range tempArr {
		fmt.Printf("%v ", val)
	}
}

func Print(arr []int, i int) {
	fmt.Printf("Printing Sorted array %v \n", i)
	for _, val := range arr {
		fmt.Printf("%v ", val)
	}
	fmt.Println("")
}

func SortArray(arr []int, i int) {
	sort.Ints(arr)
	Print(arr, i)
	wg.Done()
}

func MergeSort(arr1 []int, arr2 []int) []int {
	length := len(arr1) + len(arr2)
	var arr = make([]int, length)
	i, j := 0, 0
	k := 0
	for i, j = 0, 0; i < len(arr1) && j < len(arr2); {
		if arr1[i] < arr2[j] {
			arr[k] = arr1[i]
			k++
			i++
		} else {
			arr[k] = arr2[j]
			k++
			j++
		}
	}

	for i < len(arr1) {
		arr[k] = arr1[i]
		i++
		k++
	}

	for j < len(arr2) {
		arr[k] = arr2[j]
		j++
		k++
	}
	return arr
}
