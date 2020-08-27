package chapter4

import "fmt"

// chapter 4
type quantity int
type dozen []quantity

func display(i int) {
	fmt.Print(i)
	fmt.Print("\n")
}

// method
func (q quantity) greaterThan(i int) bool { return int(q) > i }
func (q *quantity) increment()            { *q++ }
func (q *quantity) decrement()            { *q-- }

// struct
type Item struct {
	name     string
	price    float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

type dimension struct {
	width, height, length float64
}

type Item2 struct {
	name             string
	price            float64
	quantity         int
	packageDimension dimension
	itemDimension    dimension
}

// interface
type shaper interface {
	area() float64
}

func describe(s shaper) {
	fmt.Println("area :", s.area())
}

type rect struct{ width, height float64 }

func (r rect) area() float64 {
	return r.width * r.height
}

func Run_exam() {
	var d dozen
	for i := quantity(1); i <= 12; i++ {
		d = append(d, i)
	}
	fmt.Println(d)

	var q quantity = 3
	display(int(q))

	// method
	q = quantity(3)
	q.increment()
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))
	q.decrement()
	fmt.Printf("Is q(%d) greater than %d? %t \n", q, 3, q.greaterThan(3))

	// struct
	shirt := Item{name: "Men's Slim-Fit Shirt", price: 25000, quantity: 3}
	shoes := Item{name: "Sports Shoes", price: 30000}
	dress := Item{name: "Stripe Shift Dress", quantity: 2}
	phone := Item{"Amazon Fire Phone, 32GB", 21900, 1}

	fmt.Println(shirt)
	fmt.Println(shoes)
	fmt.Println(dress)
	fmt.Println(phone)

	var t Item
	t.name = "Men's Slim-Fit Shirt"
	t.price = 25000
	t.quantity = 3

	fmt.Println(t.name)
	fmt.Println(t.price)
	fmt.Println(t.quantity)
	fmt.Println(t.Cost())

	shoes2 := Item2{
		"Sports Shoes", 30000, 2,
		dimension{30, 270, 20},
		dimension{50, 300, 30},
	}

	fmt.Printf("%#v\n", shoes2.itemDimension)
	fmt.Printf("%#v\n", shoes2.packageDimension)

	fmt.Println(shoes2.packageDimension.width)
	fmt.Println(shoes2.packageDimension.height)
	fmt.Println(shoes2.packageDimension.length)

	// interface
	r2 := rect{3, 4}
	describe(r2)
}
