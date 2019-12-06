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
	ss := strings.TrimRight(string(f), "\n")
	s := strings.Split(ss, ",")
	opcodes := []int{}
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		opcodes = append(opcodes, i)
	}

	output := []int{}
	i := 0
	for {
		// opcode: the two digits on the right
		// counting from right:
		// 3rd, 4th and possibly 5th digit: parameter codes
		// 0: as in day02, is a position mode
		// 1: immediate mode
		// 1002,4,3,4,33
		if opcodes[i] == 99 {
			break
		}

		if opcodes[i] == 0 {
			continue
		}

		// 1002,4,3,4,33
		v1, v2, v3 := 0, 0, 0
		op := ""
		sopcode := strconv.Itoa(opcodes[i])
		lenop := len(sopcode)

		if opcodes[i] > 99 {
			op = sopcode[lenop-2:]
			param1 := sopcode[lenop-3 : lenop-2]
			param2 := "0"
			//param3 := ""
			if lenop > 3 {
				param2 = sopcode[lenop-4 : lenop-3]
			}
			if param1 == "0" {
				v1 = opcodes[opcodes[i+1]]
			} else {
				v1 = opcodes[i+1]
			}
			if param2 == "0" && op != "04" {
				v2 = opcodes[opcodes[i+2]]
			} else {
				v2 = opcodes[i+2]
			}
			v3 = opcodes[i+3]

		}

		if opcodes[i] == 1 {
			sum := 0
			s1 := opcodes[i+1]
			s2 := opcodes[i+2]
			sum += opcodes[s1] + opcodes[s2]
			s3 := opcodes[i+3]
			opcodes[s3] = sum
			i += 4
		} else if op == "01" {
			opcodes[v3] = v1 + v2
			i += 4
		} else if opcodes[i] == 2 {
			sum := 0
			s1 := opcodes[i+1]
			s2 := opcodes[i+2]
			sum += opcodes[s1] * opcodes[s2]
			s3 := opcodes[i+3]
			opcodes[s3] = sum
			i += 4
		} else if op == "02" {
			opcodes[v3] = v1 * v2
			i += 4
		} else if opcodes[i] == 3 {
			// input: 1
			location := opcodes[i+1]
			opcodes[location] = 1
			i += 2
		} else if opcodes[i] == 4 {
			output = append(output, opcodes[opcodes[i+1]])
			i += 2
		} else if op == "04" {
			output = append(output, opcodes[i+1])
			i += 2
		}
	}
	fmt.Println("output: ", output)
}

// opcodes
// 1: add the numbers in the positions of the next two values
// 2: same thing, but multiply instead of adding
// 99: halt immediately
// 3: takes a single integer as input and saves it to the position given by its
//    only parameter.
//    For example, the instruction 3,50 would take an input value and store it
//    at address 50.
// 4: outputs the value of its only parameter.
// 	  For example, the instruction 4,50 would output the value at address 50.
