package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func parse(s string) int64 {
	el, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		panic("unable to parse")
	}

	return int64(el)
}

func parseInput(input *os.File) [][]int64 {
	data, err := ioutil.ReadAll(input)
	if err != nil {
		panic("Invalid input")
	}

	lines := strings.Split(string(data), "\n")

	res := [][]int64{}
	for _, line := range lines[2:] {
		if len(line) == 0 {
			continue
		}
		parsed := []int64{}
		for _, char := range strings.Split(strings.TrimSpace(line), " ") {
			if len(char) == 0 {
				continue
			}
			parsed = append(parsed, parse(char))
		}
		res = append(res, parsed)
	}

	return res
}

type matrix struct {
	m        [][]int64
	mapping  map[string]struct{}
	maxCount int64
}

func newMatrix(list [][]int64) matrix {
	m := matrix{m: list}
	m.mapping = make(map[string]struct{})
	return m
}

type pair struct {
	x, y int
}

func (p pair) s() string {
	return fmt.Sprintf("%d:%d", p.x, p.y)
}

func (m matrix) investigate(list []pair) int64 {
	if len(list) == 0 {
		return 0
	}

	p := list[0]
	list = list[1:]
	if _, ok := m.mapping[p.s()]; ok {
		return m.investigate(list)
	}

	m.mapping[p.s()] = struct{}{}

	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}
			if p.x+i >= 0 && p.x+i < len(m.m) &&
				p.y+j >= 0 && p.y+j < len(m.m[0]) &&
				m.m[p.x+i][p.y+j] == 1 {
				list = append(list, pair{p.x + i, p.y + j})
			}
		}
	}

	val := m.m[p.x][p.y]
	return val + m.investigate(list)
}

func (m matrix) getCount() int64 {
	for x := 0; x < len(m.m); x++ {
		for y := 0; y < len(m.m[0]); y++ {

			p := pair{x, y}

			list := []pair{p}

			groupCount := m.investigate(list)
			if groupCount > m.maxCount {
				m.maxCount = groupCount
			}
		}
	}
	return m.maxCount
}

func main() {
	m := newMatrix(parseInput(os.Stdin))

	count := m.getCount()

	fmt.Printf("%d\n", count)
}
