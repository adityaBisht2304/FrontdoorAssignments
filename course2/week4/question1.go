package course2week4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("Grass")
}

func (c Cow) Move() {
	fmt.Println("Walk")
}

func (c Cow) Speak() {
	fmt.Println("Moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("Worms")
}

func (b Bird) Move() {
	fmt.Println("Fly")
}

func (b Bird) Speak() {
	fmt.Println("Peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("Mice")
}

func (s Snake) Move() {
	fmt.Println("Slither")
}

func (s Snake) Speak() {
	fmt.Println("Hisss")
}

func Question_1() {
	animalMap := make(map[string]Animal)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(">")
		byteArr, _, _ := reader.ReadLine()
		stringArr := strings.Split(string(byteArr), " ")

		if stringArr[0] == "newAnimal" {
			_, ok := animalMap[stringArr[1]]
			if ok {
				fmt.Println("Animal is already created with the provided name")
				continue
			}
			switch stringArr[2] {
			case "cow":
				animalMap[stringArr[1]] = Cow{}
			case "bird":
				animalMap[stringArr[1]] = Bird{}
			case "snake":
				animalMap[stringArr[1]] = Snake{}
			}
			fmt.Println("Created new animal")
		} else if stringArr[0] == "query" {
			val, ok := animalMap[stringArr[1]]
			if ok {
				switch stringArr[2] {
				case "eat":
					val.Eat()
				case "move":
					val.Move()
				case "speak":
					val.Speak()
				}
			} else {
				fmt.Println("Animal with provided name does not exist")
			}
		}

	}
}
