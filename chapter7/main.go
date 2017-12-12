package main

import ("fmt"; "math")

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1

	return math.Sqrt(a*a + b*b)
}

// Shape a geometric shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

// Circle a simple circle shape
type Circle struct {
	x, y, r float64
}
// Area calculate the area of a circle
func (c *Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

// Perimeter return the circumfurence of the Circle
func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}


// Rectangle a shape defined by 2 points
type Rectangle struct {
	x1, y1, x2, y2 float64
}
// Area return the area of the rectangle shape
func (r *Rectangle) Area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1) 
	return l * w
	
}

// Perimeter teturn the perimeter of the Rectangle
func (r *Rectangle) Perimeter() float64 {
	return 2*(math.Abs(r.x1 - r.x2) + math.Abs(r.y1 - r.y2))
}

// MultiShape collection of Shapes
type MultiShape struct {
	shapes []Shape
}

// Area return the combined areas of all the Shapes
func (m *MultiShape) Area() float64 {
	var area float64
	for _,s := range m.shapes {
		area += s.Area()
	}
	return area
}

// Perimeter return the combined perimeters of all the Shapes
func (m *MultiShape) Perimeter() float64 {
	var area float64
	for _,s := range m.shapes {
		area += s.Perimeter()
	}
	return area
}


func main() {

	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}

	fmt.Println(r.Area())
	fmt.Println(c.Area())
	fmt.Println(totalArea(&c, &r))

	multiShape := MultiShape{
		shapes: []Shape{
			&Circle{0,0,5},
			&Rectangle{0,0,10,10},
		},
	}

	fmt.Println(multiShape.Area())
}