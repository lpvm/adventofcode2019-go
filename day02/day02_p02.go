package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			var tryCodes []int
			tryCodes = make([]int, len(opcodes))
			copy(tryCodes, opcodes)
			tryCodes[1] = noun
			tryCodes[2] = verb

			i := 0
			for {

				if tryCodes[i] == 99 {
					break
				}

				if tryCodes[i] == 1 {
					sum := 0
					s1 := tryCodes[i+1]
					s2 := tryCodes[i+2]
					sum += tryCodes[s1] + tryCodes[s2]
					s3 := tryCodes[i+3]
					tryCodes[s3] = sum
				} else if tryCodes[i] == 2 {
					sum := 0
					s1 := tryCodes[i+1]
					s2 := tryCodes[i+2]
					sum += tryCodes[s1] * tryCodes[s2]
					s3 := tryCodes[i+3]
					tryCodes[s3] = sum
				}
				i += 4
				if tryCodes[0] == 19690720 {
					fmt.Println("noun: ", noun, "  verb: ", verb, "  100 * noun + verb: ", 100*noun+verb)
					os.Exit(0)
				}
			}

		}
	}
}

// opcodes
// 1: add the numbers in the positions of the next two values
// 2: same thing, but multiply instead of adding
// 99: halt immediately
