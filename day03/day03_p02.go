package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// start point [0, 0]
// UP: [0, +]
// DOWN: [0, -]
// RIGHT: [+, 0]
// LEFT: [-, 0]

func seq(min, max int) []int {
	s := []int{}
	if max > min {
		for i := min + 1; i <= max; i++ {
			s = append(s, i)
		}
	} else {
		for i := min - 1; i >= max; i-- {
			s = append(s, i)
		}
	}
	return s
}

// pos 0 -> change x coordinate
// pos 1 -> change y coordinate
func segment(seq []int, pos int, val int) [][]int {
	segm := [][]int{}
	switch pos {
	case 0:
		for _, k := range seq {
			p := []int{k, val}
			segm = append(segm, p)
		}
	case 1:
		for _, k := range seq {
			p := []int{val, k}
			segm = append(segm, p)
		}
	}

	return segm
}

func path(wire []string, c chan [][]int) {
	p := [][]int{}
	//curr := []int{}
	local := [2]int{0, 0}
	dir := ""
	for _, k := range wire {
		dir = string(rune(k[0]))
		way, err := strconv.Atoi(k[1:])
		//fmt.Println("dir: ", dir, "  way: ", way)
		if err != nil {
			log.Fatal(err)
		}
		switch dir {
		case "U":
			s := seq(local[1], local[1]+way)
			//fmt.Println("Sequence s: ", s)
			local[1] = local[1] + way
			segm := segment(s, 1, local[0])
			//fmt.Println("Segments: ", segm)
			p = append(p, segm...)
			//fmt.Println("UP k: ", k, "  segm: ", segm, "  p: ", p)
		case "D":
			s := seq(local[1], local[1]-way)
			//fmt.Println("Sequence s: ", s)
			local[1] = local[1] - way
			segm := segment(s, 1, local[0])
			//fmt.Println("Segments: ", segm)
			p = append(p, segm...)
			//fmt.Println("UP k: ", k, "  segm: ", segm, "  p: ", p)
		case "L":
			s := seq(local[0], local[0]-way)
			//fmt.Println("Sequence s: ", s)
			local[0] = local[0] - way
			segm := segment(s, 0, local[1])
			//fmt.Println("Segments: ", segm)
			p = append(p, segm...)
		case "R":
			s := seq(local[0], local[0]+way)
			//fmt.Println("Sequence s: ", s)
			local[0] = local[0] + way
			segm := segment(s, 0, local[1])
			//fmt.Println("Segments: ", segm)
			p = append(p, segm...)
		}
	}
	c <- p
}

func sort(unsorted [][]int) {
	if len(unsorted) <= 1 {
		return
	}

	//fmt.Println("unsorted: BEF ", unsorted)
	isUnsorted := false
	idx := 0
	absp00 := 0
	absp01 := 0
	absp10 := 0
	absp11 := 0
	for {
		//fmt.Println("idx: ", idx)
		if idx < len(unsorted)-1 {
			p0 := unsorted[idx]
			p1 := unsorted[idx+1]
			//fmt.Println("Analysing: p0 and p1: ", p0, "  ", p1)
			if p0[0] < 0 {
				absp00 = -p0[0]
			} else {
				absp00 = p0[0]
			}
			if p0[1] < 0 {
				absp01 = -p0[1]
			} else {
				absp01 = p0[1]
			}
			if p1[0] < 0 {
				absp10 = -p1[0]
			} else {
				absp10 = p1[0]
			}
			if p1[1] < 0 {
				absp11 = -p1[1]
			} else {
				absp11 = p1[1]
			}
			if absp10+absp11 < absp00+absp01 {
				unsorted[idx] = p1
				unsorted[idx+1] = p0
				//fmt.Println("change p0 and p1: ", p0, "  ", p1)
				isUnsorted = true
			}
		}
		if idx == len(unsorted)-2 && !isUnsorted {
			break
		} else if idx == len(unsorted)-2 {
			//time.Sleep(1 * time.Second)
			idx = -1
			isUnsorted = false
		}
		//fmt.Println("unsorted: AFT ", unsorted)
		idx++
	}
}

func convertMap(p [][]int) map[[2]int]int {
	a := [2]int{}
	m := make(map[[2]int]int)
	for _, k := range p {
		a[0] = k[0]
		a[1] = k[1]
		m[a] = 0
	}
	return m
}

func findNearest(m map[[2]int]int, l [][]int) int {
	d := 0
	a := [2]int{}
	for _, k := range l {
		a[0] = k[0]
		a[1] = k[1]
		if _, ok := m[a]; ok == true {
			//fmt.Println("d: ", d, "  k: ", k, "   a: ", a)
			//fmt.Println("map ", m)
			if k[0] < 0 {
				d += -k[0]
			} else {
				d += k[0]
			}
			if k[1] < 0 {
				d += -k[1]
			} else {
				d += k[1]
			}
			break
		}
	}
	return d
}

func findShortest(p0 [][]int, p1 [][]int, m map[[2]int]int) int {
	all := [][]int{}
	a := [2]int{}
	for _, k := range p1 {
		a[0] = k[0]
		a[1] = k[1]
		if _, ok := m[a]; ok == true {
			all = append(all, k)
		}
	}

	idx0 := []int{}
	idx1 := []int{}
	for _, k := range all {
		for i0, k0 := range p0 {
			if k[0] == k0[0] && k[1] == k0[1] {
				idx0 = append(idx0, i0+1)
				break
			}
		}
		for i1, k1 := range p1 {
			if k[0] == k1[0] && k[1] == k1[1] {
				idx1 = append(idx1, i1+1)
				break
			}
		}
	}
	stepsSl := []int{}
	min := 99999999999
	for i := 0; i < len(idx0); i++ {
		sum := idx0[i] + idx1[i]
		stepsSl = append(stepsSl, sum)
		if sum < min {
			min = sum
		}
	}
	return min
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	st := strings.Split(string(file), "\n")
	wire0 := strings.Split(st[0], ",")
	wire1 := strings.Split(st[1], ",")
	st = nil
	//fmt.Println("wire0: ", wire0)
	c0 := make(chan [][]int)
	c1 := make(chan [][]int)
	go path(wire0, c0)
	go path(wire1, c1)
	path0 := <-c0
	path1 := <-c1
	//fmt.Println("path0 ", path0)
	mpath0 := convertMap(path0)
	//fmt.Println("path1 ", path1)
	path1Sorted := make([][]int, len(path1))
	copy(path1Sorted, path1)
	sort(path1Sorted)
	//fmt.Println("path1 sorted: ", path1)
	nearest := findNearest(mpath0, path1Sorted)
	fmt.Println("nearest ", nearest)
	shortest := findShortest(path0, path1, mpath0)
	fmt.Println("shortest ", shortest)
}
