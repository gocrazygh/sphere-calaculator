//Sphere Calculator by GOCrazy
//This is my first code
//It will take a radius of a sphere and will calculate
//and show it's surface area, volume, and diameter

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func sfa(f float64) float64 {
	return math.Pi * math.Pow(f, 2) * 4
}

func vol(f float64) float64 {
	return math.Pi * math.Pow(f, 3) * math.Pi
}

func dim(f float64) float64 {
	return f * 2
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a sphere's radius: ")
	s.Scan()
	f, _ := strconv.ParseFloat(s.Text(), 64)

	sfaMain := sfa(f)
	volMain := vol(f)
	dimMain := dim(f)

	fmt.Printf("The surfaace area of the sphere: %.3f \n", sfaMain)
	fmt.Printf("The volume of the sphere: %.3f \n", volMain)
	fmt.Printf("The diameter of the sphere: %.3f \n", dimMain)
}
