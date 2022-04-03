package main

import (
	"bufio"
	"fmt"
	"os"
)

type Sns struct {
	followers [][]int
	n         int
}

type SnsRestore interface {
	RestoreFollowers()
	Print()
}

func (s *Sns) RestoreFollowers(q int) error {
	var command, a, b int

	log := bufio.NewReader(os.Stdin)

	for l := 0; l < q; l++ {
		fmt.Fscan(log, &command)

		if command == 1 {
			fmt.Fscan(log, &a, &b)
			a--
			b--

			s.followers[a][b] = 1
		} else if command == 2 {
			fmt.Fscan(log, &a)
			a--

			for j := 0; j < s.n; j++ {
				if s.followers[j][a] == 1 {
					s.followers[a][j] = 1
				}
			}
		} else if command == 3 {
			fmt.Fscan(log, &a)
			a--

			follow_list := []int{}
			for j := 0; j < s.n; j++ {
				if s.followers[a][j] == 1 {
					follow_list = append(follow_list, j)
				}
			}

			for _, i := range follow_list {
				for j := range s.followers[i] {
					if s.followers[i][j] == 1 && a != j {
						s.followers[a][j] = 1
					}
				}
			}
		} else {
			return fmt.Errorf("unkown command '%d'", command)
		}
	}

	return nil
}

func (s *Sns) Print() {
	result := bufio.NewWriter(os.Stdout)
	defer result.Flush()

	for i := 0; i < s.n; i++ {
		for j := 0; j < s.n; j++ {
			if s.followers[i][j] == 1 {
				fmt.Fprintf(result, "Y")
			} else {
				fmt.Fprintf(result, "N")
			}
		}
		fmt.Fprintf(result, "\n")
	}
}

func main() {
	var n, q int

	_, e := fmt.Scanf("%d %d", &n, &q)
	if e != nil {
		fmt.Println(e)
	}

	s := &Sns{
		followers: make([][]int, n),
		n:         n,
	}
	for i := 0; i < n; i++ {
		s.followers[i] = make([]int, n)
	}

	e = s.RestoreFollowers(q)
	if e != nil {
		fmt.Println(e)
		return
	}
	s.Print()
}
