package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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
		sum = sum + k/3 - 2
	}
	fmt.Println(sum)
}
