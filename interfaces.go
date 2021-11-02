package main
import ("fmt"
		"math"	
)

type geometry interfaces {
	area() float64
	perim() float64
}
type rect struct{
	width,height float64
}
type circle struct{
	radius float64
}
func (r rect) area() float64{
	return r.width * r.height
}

func (c circle) area() float64{
	return math.pi * c.radius * c.radius
}

func (r rect) perim() float64{
	return 2*r.width + 2*r.height
}

func (c circle) perim() float64{
	return 2*math.pi * c.radius
}

func measure(g geometry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){

	r :=rect{width:3, height :4}
	c := circle {radius : 5}
	measure(r)

}