package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(string(f), ",")
	opcodes := []int{}
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		opcodes = append(opcodes, i)
	}

	opcodes[1] = 12
	opcodes[2] = 2
	i := 0
	for {

		if opcodes[i] == 99 {
			fmt.Println("opcodes[0]: ", opcodes[0])
			break
		}

		if opcodes[i] == 1 {
			sum := 0
			s1 := opcodes[i+1]
			s2 := opcodes[i+2]
			sum += opcodes[s1] + opcodes[s2]
			s3 := opcodes[i+3]
			opcodes[s3] = sum
		} else if opcodes[i] == 2 {
			sum := 0
			s1 := opcodes[i+1]
			s2 := opcodes[i+2]
			sum += opcodes[s1] * opcodes[s2]
			s3 := opcodes[i+3]
			opcodes[s3] = sum
		}
		i += 4
	}
	fmt.Println("opcodes: ", opcodes)
	fmt.Println("len(opcodes): ", len(opcodes))
}

// opcodes
// 1: add the numbers in the positions of the next two values
// 2: same thing, but multiply instead of adding
// 99: halt immediately
