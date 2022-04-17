package course2week2

import (
	"fmt"
	"math"
)

func GenDisplacement(acc float64, vel float64, initDis float64) func(float64) float64 {
	fn := func(time float64) float64 {
		s := ((acc / 2.0) * math.Pow(time, 2)) + (vel * time) + initDis

		return s
	}
	return fn
}

func Question_1() {
	var acc float64
	var vel float64
	var initDis float64

	fmt.Printf("Enter Acceleration\n")
	fmt.Scan(&acc)
	fmt.Printf("Enter Velocity\n")
	fmt.Scan(&vel)
	fmt.Printf("Enter Initial Displacement\n")
	fmt.Scan(&initDis)

	fn := GenDisplacement(acc, vel, initDis)

	var time float64
	fmt.Printf("Enter Time\n")
	fmt.Scan(&time)
	result := fn(time)
	fmt.Printf("Displacement : %v\n", result)

	fmt.Printf("Enter Time\n")
	fmt.Scan(&time)
	result = fn(time)
	fmt.Printf("Displacement : %v\n", result)
}
