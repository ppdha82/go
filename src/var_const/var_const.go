package var_const

import "fmt"

func RunExam() {
	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)

	fmt.Println(Apple)
	fmt.Println(Grape)
	fmt.Println(Orange)
}
