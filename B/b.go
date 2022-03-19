package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var days int
	_, err := fmt.Scanf("%d", &days)
	if err != nil {
		fmt.Println(err)
	}

	sales := make([]int, days)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < days; i++ {
		scanner.Scan()
		sales[i], err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
	}

	for i := 0; i < days-1; i++ {
		i1 := i + 1
		if sales[i1] == sales[i] {
			fmt.Println("stay")
		} else if sales[i1] < sales[i] {
			fmt.Printf("down %d\n", sales[i]-sales[i1])
		} else {
			fmt.Printf("up %d\n", sales[i1]-sales[i])
		}
	}
}
