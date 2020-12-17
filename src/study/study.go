package study

import "fmt"

func RunExam() {
	var_const()
	condition(0, 10)
	condition(1, 10)
	condition(1, 0)
	switch_case(0)
	switch_case(2)
	switch_case(3)
	for_loop()
}

func normal_for_loop() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
}

func only_condition_for_loop() {
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println("n = ", n)
}

func range_for_loop() {
	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		fmt.Println("index = ", index, "> name = ", name)
	}
}

func etc_for_loop() {
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			fmt.Println("a == 5 but a += a => ", a)
			continue
		}
		a++
		if a < 10 {
			fmt.Println("a = ", a)
			break
		}
	}
	if a == 11 {
		fmt.Println("a = ", a, "and go to END")
		goto END
	}
	fmt.Println("a = ", a)

END:
	fmt.Println("This is END")
}

func break_label_for_loop() {
	i := 0
L1:
	for {
		if i == 0 {
			fmt.Println("I met break L1")
			break L1
		}
	}
	fmt.Println("Just OK")
}

func for_loop() {
	// 일반 for 문
	normal_for_loop()
	// 조건식만 쓰는 for 문
	only_condition_for_loop()
	// for range 문 (like foreach)
	range_for_loop()
	// break, continue, goto 문을 이용한 for 문
	etc_for_loop()
	// break label 을 이용한 for 문
	break_label_for_loop()
}

func switch_case(category int) {
	var name string
	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3:
		name = "Blog"
	default:
		name = "other"
	}
	fmt.Println("name = ", name)
}

func condition(k int, t int) {
	if k = 1; t != 0 {
		fmt.Println("k is ", k, ". t is not 0\n")
	} else {
		fmt.Println("k is ", k, "t is ", t)
	}
}

func var_const() {
	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)

	fmt.Println(Apple)
	fmt.Println(Grape)
	fmt.Println(Orange)
}
