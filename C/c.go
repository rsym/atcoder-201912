package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	value int
}

type DataArray []Data

func (d DataArray) Len() int {
	return len(d)
}

func (d DataArray) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d DataArray) Less(i, j int) bool {
	return d[i].value < d[j].value
}

func main() {
	d := make([]Data, 0)

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := strings.Split(s.Text(), " ")

	l := len(input)
	for i := 0; i < l; i++ {
		t, e := strconv.Atoi(input[i])
		if e != nil {
			fmt.Println(e)
			return
		}
		d = append(d, Data{value: t})
	}

	sort.Sort(DataArray(d))
	fmt.Println(d[l-3].value)
}
