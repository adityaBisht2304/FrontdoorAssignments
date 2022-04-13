package course1week4

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Question_1() {
	nameAddr := make(map[string]string, 1)

	name := ""
	addr := ""
	fmt.Printf("Enter Name\n")
	fmt.Scan(&name)
	fmt.Printf("Enter Address\n")
	fmt.Scan(&addr)
	nameAddr[name] = addr

	byteArr, err := json.Marshal(nameAddr)
	if err != nil {
		fmt.Printf("Error in marshalling\n")
		return
	}

	fmt.Printf("Printing JSON Object\n")
	fmt.Printf(string(byteArr))
	fmt.Println()
}

type Name struct {
	fName string
	lName string
}

func Question_2() {
	fh, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("Error in opening file\n")
	}
	defer fh.Close()

	reader := bufio.NewReader(fh)

	byteArr, _, err := reader.ReadLine()

	structArr := make([]Name, 0)

	for err == nil {
		stringArr := strings.Split(string(byteArr), " ")
		structArr = append(structArr, Name{
			fName: stringArr[0],
			lName: stringArr[1],
		})
		byteArr, _, err = reader.ReadLine()
	}

	for index, val := range structArr {
		fmt.Printf("Structure %v : %v\n", index, val)
	}

}
