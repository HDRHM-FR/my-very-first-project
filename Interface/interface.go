package Interface

import "fmt"

type geo interface {
	area() float64
}

type squar struct {
	length float64
}

func (s *squar) area() float64 {
	return s.length * 4
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return (c.radius * c.radius) * 3.14
}

func printArea(g geo) {

	fmt.Print("Area ")

	switch g.(type) {
	case *squar:
		fmt.Printf("squar : %v\n", g.area())
	case *circle:
		fmt.Printf("circle : %v\n", g.area())
	}
}

func Salam() {

	circle := circle{5}
	square := squar{4}
	printArea(&circle)
	printArea(&square)

}
