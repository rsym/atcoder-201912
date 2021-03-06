package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Sequence struct {
	data    []int
	counter []int
	size    int
}

type SequenceChecker interface {
	Check()
}

func (d *Sequence) Check() (int, int, error) {
	for _, v := range d.data {
		d.counter[v]++
	}

	var x, y int
	for i, v := range d.counter {
		if v == 0 {
			x = i
		} else if v > 1 {
			y = i
		}

		if x != 0 && y != 0 {
			return x, y, errors.New("Incorrect")
		}
	}

	return 0, 0, nil
}

func main() {
	var n int
	_, e := fmt.Scanf("%d", &n)
	if e != nil {
		fmt.Println(e)
	}

	d := &Sequence{
		data:    make([]int, n+1),
		counter: make([]int, n+1),
		size:    n + 1,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i, _ := range d.data[1:] {
		scanner.Scan()
		d.data[i], e = strconv.Atoi(scanner.Text())
		if e != nil {
			fmt.Println(e)
		}
	}

	x, y, e := d.Check()
	if e != nil {
		fmt.Printf("%d %d\n", y, x)
	} else {
		fmt.Println("Correct")
	}
}
