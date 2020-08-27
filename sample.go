package main

import (
	"fmt"
	"image/color"
	test_lib "lib"
	"strings"
	"unicode"
	// chapter 4
)

func f1() {
	fmt.Println("f1 - start")
	defer f2()
	fmt.Println("f1 - end")
}

func f2() {
	fmt.Println("f2 - deferred")
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func makeSuffix(suffix string) func(string) string {
	fmt.Println("RUN makeSuffix suffix = ", suffix)
	return func(name string) string {
		fmt.Println("suffix = ", suffix, ";name = ", name)
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

var v rune

func init() {
	v = '1'
}

type rect1 struct {
	x0, y0, x1, y1 int
	color.RGBA
}

func resize(rect *rect1, width, height int) {
	(*rect).x1 += width
	rect.y1 += height
}

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

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2, i := 0, 0
	for i < 10 {
		sum2 += i
		i++
	}
	fmt.Println(sum)

	type Rect struct {
		width  int
		height int
	}

	//r := Rect{1, 2}; fmt.Println(r.width * 2 + r.height * 2)
	r := Rect{1, 2}
	fmt.Println(r.width*2 + r.height*2)

	f1()
	f()

	fmt.Println("\nmakeSuffix - start")
	addZip := makeSuffix(".zip")
	addTgz := makeSuffix(".tar.gz")
	fmt.Println("makeSuffix - end")

	fmt.Println(addTgz("go1.5.1.src"))
	fmt.Println(addZip("go1.5.1.windows-amd64"))

	f := func(c rune) bool {
		return unicode.Is(unicode.Hangul, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 월드", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))

	fmt.Println(test_lib.IsDigit('1'))
	fmt.Println(test_lib.IsDigit('a'))
	fmt.Println(test_lib.IsDigit('\t'))

	fmt.Println("2.5.4 init() 함수 - Start")
	fmt.Println(test_lib.IsDigit(v))
	fmt.Println("2.5.4 init() 함수 - End")

	s := "hello"
	for j := 0; j < len(s); j++ {
		fmt.Printf("j[%d] = %c\n", j, s[j])
	}
	fmt.Println(s[1:4])

	var a [5]int

	fmt.Println("len(a) = ", len(a))

	numbers := []int{3, 4, 5, 7, 8, 4, 6, 8, 32, 4}
	for i, v := range numbers {
		fmt.Println(i, v)
	}

	sum = 0
	for _, v := range numbers {
		sum += v
		fmt.Println("sum = ", sum)
	}
	fmt.Println("new sum")

	sum = 0
	for v := range numbers {
		fmt.Println("v = ", v)
	}

	b := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(b, "=", b[:3], b[3:5], b[5:])

	ns1 := []int{1, 2, 3}
	ns2 := []int{5, 6, 7}
	ns3 := []int{8, 9, 10, 11}

	ns1 = append(ns1, 4, 5)
	fmt.Println(ns1)
	ns1 = append(ns1, ns2...)
	fmt.Println(ns1)
	ns1 = append(ns1, ns3[1:3]...)

	fmt.Println(ns1)

	k := make([]int, 0, 3)
	for i := 0; i < 9; i++ {
		k = append(k, i)
		fmt.Printf("len: %d, cap: %d, %v\n", len(k), cap(k), k)
	}

	numberMap := map[string]int{}
	numberMap["one"] = 1
	numberMap["two"] = 2
	numberMap["three"] = 3

	fmt.Println(numberMap)

	numberMap1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(numberMap1)

	numberMap2 := make(map[string]int, 3)
	numberMap2["one"] = 1
	numberMap2["two"] = 2
	numberMap2["three"] = 3
	fmt.Println(numberMap2)
	fmt.Println("for loop")
	for l, v := range numberMap2 {
		fmt.Println(l, v)
	}

	fmt.Println(numberMap["one"])
	fmt.Println(numberMap["zero"])

	find_key := "zero"
	if v, ok := numberMap[find_key]; ok {
		fmt.Println(find_key, "is in numberMap. value: ", v)
	} else {
		fmt.Println(find_key, "is not in numberMap")
	}

	find_key2 := "four"
	if _, ok := numberMap[find_key2]; !ok {
		numberMap[find_key2] = 4
		fmt.Println(find_key2, "is not in numberMap. value: 4")
	}
	fmt.Println(numberMap)

	rect2 := rect1{2, 4, 10, 20, color.RGBA{0xff, 0, 0, 0xff}}
	resize(&rect2, 10, 10)
	fmt.Println(rect2)

	// chapter 4
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
