package chap3

import (
	"fmt"
	"image/color"
)

type rect1 struct {
	x0, y0, x1, y1 int
	color.RGBA
}

func resize(rect *rect1, width, height int) {
	(*rect).x1 += width
	rect.y1 += height
}

func Run_exam() {
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

	sum := 0
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
}
