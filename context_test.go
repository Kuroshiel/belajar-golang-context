package belajargolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Membuat Context

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

// Context With Value

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	// Context Get Value

	fmt.Println(contextF.Value("f")) // dapat
	fmt.Println(contextF.Value("c")) // dapat milik parent
	fmt.Println(contextF.Value("b")) // tidak dapat, beda parent

	fmt.Println(contextA.Value("b")) // tidak bisa menngambil data child
}

// Context With Cencel

// Goroutine Leak

func CreateCounterLeak() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {

			destination <- counter
			counter++

		}
	}()

	return destination
}

func TestGoroutineLeak(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	destination := CreateCounterLeak()

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

// Contex With Cencel, signal cencel

func CreateCounterWithCencel(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCencel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cencel := context.WithCancel(parent)

	destination := CreateCounterWithCencel(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}
	cencel() // menigirim signal cencel ke context

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

// Context With Timeout

func CreateCounterWithTimeout(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulai Slow
			}
		}
	}()

	return destination
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cencel := context.WithTimeout(parent, 5*time.Second)
	defer cencel()

	destination := CreateCounterWithTimeout(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
