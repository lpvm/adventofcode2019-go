package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func rec(t int) int {
	f := 0
	for {
		t = t/3 - 2
		if t <= 0 {
			break
		}
		f += t
	}
	return f
}

func main() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(file), "\n")
	inpInt := []int{}
	for _, k := range input {
		i, err := strconv.Atoi(k)
		if err != nil {
			continue
		}
		inpInt = append(inpInt, i)
	}

	sum := 0
	for _, k := range inpInt {
		module := k/3 - 2
		sum += module
		sum += rec(module)

	}

	fmt.Println(sum)
}
