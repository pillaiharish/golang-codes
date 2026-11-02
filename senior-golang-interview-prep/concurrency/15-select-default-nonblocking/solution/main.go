package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	dropped := 0
	for _, value := range []int{10, 20} {
		select {
		case ch <- value:
			fmt.Println("accepted:", value)
		default:
			dropped++
			fmt.Println("dropped:", value)
		}
	}
	select {
	case v := <-ch:
		fmt.Println("received:", v)
	default:
		fmt.Println("nothing ready")
	}
	select {
	case v := <-ch:
		fmt.Println("received:", v)
	default:
		fmt.Println("channel empty")
	}
	fmt.Println("drop count:", dropped)
}
