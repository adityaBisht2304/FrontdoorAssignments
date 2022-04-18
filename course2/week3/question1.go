package course2week3

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Printf("%v\n", a.food)
}

func (a Animal) Move() {
	fmt.Printf("%v\n", a.locomotion)
}

func (a Animal) Speak() {
	fmt.Printf("%v\n", a.noise)
}

func (a Animal) PrintFunction(str string) {
	switch str {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

func Question_1() {
	cow := Animal{
		food:       "Grass",
		locomotion: "Walk",
		noise:      "Moo",
	}

	bird := Animal{
		food:       "Worms",
		locomotion: "Fly",
		noise:      "Peep",
	}

	snake := Animal{
		food:       "Mice",
		locomotion: "Slither",
		noise:      "Hiss",
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(">")
		byteArr, _, _ := reader.ReadLine()
		stringArr := strings.Split(string(byteArr), " ")

		switch stringArr[0] {
		case "cow":
			cow.PrintFunction(stringArr[1])
		case "bird":
			bird.PrintFunction(stringArr[1])
		case "snake":
			snake.PrintFunction(stringArr[1])
		}
	}
}
