package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Print(err)
		return
	}

	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("%d\n", i*2)
}
