package main

type Shape interface {
	Area() float64
}

type circle struct {
	radius float64
}

func (c circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	circle := circle{radius: 5}
	rectangle := rectangle{width: 4, height: 6}

	println("Area of circle:", circle.Area())
	println("Area of rectangle:", rectangle.Area())
}
