package chapter5

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func RunExam() {
	//goRoutine()
	//goChan()
	goMutex()
	goOnce()
	goWait()
}

type counter struct {
	i int64
}

const initialValue = -500

type counter2 struct {
	i    int64
	mu   sync.Mutex
	once sync.Once
}

type counter3 struct {
	i  int64
	mu sync.Mutex
}

func goRoutine() {
	fmt.Println("go_routine 함수 시작", time.Now())

	go long()
	go short()

	time.Sleep(5 * time.Second)
	fmt.Println("go_routine 함수 종료", time.Now())
}

func goChan() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	go func() { c <- 3 }()
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func goFibonacci() {
	c2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c2)
		}
		quit <- 0
	}()

	fibonacci(c2, quit)
}

func goMutex() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	count := counter{i: 0}
	done := make(chan struct{})

	for i := 0; i < 1000; i++ {
		go func() {
			count.increment()
			done <- struct{}{}
		}()
	}
	for i := 0; i < 1000; i++ {
		<-done
	}

	fmt.Println("goMutex() - Start")
	count.display()
	fmt.Println("goMutex() - End")
}

func goOnce() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter2{i: 0}
	done := make(chan struct{})

	for i := 0; i < 1000; i++ {
		go func() {
			c.increment()
			done <- struct{}{}
		}()
	}

	for i := 0; i < 1000; i++ {
		<-done
	}

	fmt.Println("goOnce() - Start")
	c.display()
	fmt.Println("goOnce() - End")
}

func goWait() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := counter3{i: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}

	wg.Wait()

	fmt.Println("goWait() - Start")
	c.display()
	fmt.Println("goWait() - End")
}

func long() {
	fmt.Println("long 함수 시작", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("long 함수 종료", time.Now())
}

func short() {
	fmt.Println("short 함수 시작", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("short 함수 종료", time.Now())
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func (count *counter) increment() {
	count.i += 1
}

func (count *counter) display() {
	fmt.Println(count.i)
}

func (count *counter2) increment() {
	count.once.Do(func() {
		count.i = initialValue
	})
	count.mu.Lock()
	count.i += 1
	count.mu.Unlock()
}

func (count *counter2) display() {
	fmt.Println(count.i)
}

func (count *counter3) increment() {
	count.mu.Lock()
	count.i += 1
	count.mu.Unlock()
}

func (count *counter3) display() {
	fmt.Println(count.i)
}
