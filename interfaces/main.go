package main

import "fmt"

// type paymenter interface{
// 	pay(amount float32)
// }

// type payment struct{
// 	gateway paymenter
// }

// func (p payment) makePayment(amount float32){
// 	p.gateway.pay(amount)
// }

// type eSewa struct{}

// func (eS eSewa) pay(amount float32){
// 	fmt.Println("Making payment of", amount,"using eSewa.")
// }

// func main()  {
// 	es := eSewa{}
// 	py := payment{
// 		gateway: es,
// 	}
// 	py.makePayment(200)
// }

type shaper interface{
	Area() float64
}

func getArea(s shaper) {
	fmt.Println("Area is:", s.Area())
}

type square struct{
	length float64
}

func (s square) Area() float64{
	return s.length *s.length
}

type circle struct{
	radius float64
}

func (c circle) Area() float64 {
	return 3.14 * c.radius *c.radius;
}

func main()  {
	sq := square{length: 10}
	c := circle{radius: 10}

	getArea(sq)
	getArea(c)
}