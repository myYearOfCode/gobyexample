package main
import "fmt"

func main () {

	// the most basic type with a single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// a classic initial / condition / after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// this will run repeatedly until you break out of the loop
	for {
		fmt.Println("loop")
		break
	}

	// continue is like an invert break where you simply skip to the next iteration
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}