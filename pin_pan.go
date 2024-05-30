package main

import "fmt"

func main() {

	for end := 0; end <= 100; end++ {
		if end%3 == 0 {
			fmt.Println(end, "Pin")
		} else if end%5 == 0 {
			fmt.Println(end, "Pan")
		}

	}
}
