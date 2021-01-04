package study

import "fmt"

// RunExam is example code
func RunExam() {
	varConst()
	condition(0, 10)
	condition(1, 10)
	condition(1, 0)
	switchCase(0)
	switchCase(2)
	switchCase(3)
	forLoop()
	funcTest()
	anonyFuncTest()
	// closure
	runClosure()
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func runClosure() {
	next := nextValue()

	fmt.Println("[next value]")
	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())

	fmt.Println("[anotherNext value]")
	anotherNext := nextValue()
	fmt.Println(anotherNext())
	fmt.Println(anotherNext())
}

func anonyFuncTest() {
	anonymousFuncTest()
	firstFuncTest()
}

func firstFuncTest() {
	add := func(i int, j int) int {
		return i + j
	}
	r1 := calc(add, 10, 20)
	fmt.Println("r1 = ", r1)

	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	fmt.Println("r2 = ", r2)
}

type calculator func(int, int) int

func calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

func anonymousFuncTest() {
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("result = ", result)
}

func funcTest() {
	msg := "Hello"
	say(msg)
	saying(&msg)
	fmt.Println("msg = ", msg)
	talk("This", "is", "a", "book")
	talk("Hi")
	total, count := sum(1, 3, 5, 7, 9)
	fmt.Println("total = ", total, "; count = ", count)
	sums(1, 5, 10)
}

func sums(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	fmt.Println("count = ", count, "total = ", total)

	return
}

func sum(nums ...int) (int, int) {
	s := 0
	count := 0
	for _, n := range nums {
		s += n
		count++
	}
	return s, count
}

func talk(msg ...string) {
	for _, s := range msg {
		fmt.Println("s = ", s)
	}
}

func saying(msg *string) {
	fmt.Println("msging = ", *msg)
	*msg = "Changed"
}

func say(msg string) {
	fmt.Println("msg = ", msg)
}

func normalForLoop() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)
}

func onlyConditionForLoop() {
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println("n = ", n)
}

func rangeForLoop() {
	names := []string{"홍길동", "이순신", "강감찬"}

	for index, name := range names {
		fmt.Println("index = ", index, "> name = ", name)
	}
}

func etcForLoop() {
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

func breakLabelForLoop() {
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

func forLoop() {
	// 일반 for 문
	normalForLoop()
	// 조건식만 쓰는 for 문
	onlyConditionForLoop()
	// for range 문 (like foreach)
	rangeForLoop()
	// break, continue, goto 문을 이용한 for 문
	etcForLoop()
	// break label 을 이용한 for 문
	breakLabelForLoop()
}

func switchCase(category int) {
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
		fmt.Println("k is ", k, ". t is not 0")
	} else {
		fmt.Println("k is ", k, "t is ", t)
	}
}

func varConst() {
	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)

	fmt.Println(Apple)
	fmt.Println(Grape)
	fmt.Println(Orange)
}
