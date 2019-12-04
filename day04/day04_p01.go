package main

import (
	"fmt"
	"log"
	"strconv"
)

func double(p []rune) bool {
	twoadj := false
	if p[0] == p[1] || p[1] == p[2] || p[2] == p[3] ||
		p[3] == p[4] || p[4] == p[5] {
		twoadj = true
	}
	return twoadj
}

func noDecreasing(p []rune) bool {
	nodec := false
	if p[0] <= p[1] && p[1] <= p[2] && p[2] <= p[3] &&
		p[3] <= p[4] && p[4] <= p[5] {
		nodec = true
	}
	return nodec
}

func inLimits(p []rune) bool {
	inlim := false
	pass, err := strconv.Atoi(string(p))
	if err != nil {
		log.Fatal(err)
	}
	if pass >= 183564 && pass <= 657474 {
		inlim = true
	}
	// Your puzzle input is 183564-657474
	return inlim
}

func main() {
	f := "188888"
	first := []rune(f)
	count := 0
	for i0 := first[0]; i0 <= rune(54); i0++ {
		for i1 := i0; i1 <= rune(57); i1++ {
			for i2 := i1; i2 <= rune(57); i2++ {
				for i3 := i2; i3 <= rune(57); i3++ {
					for i4 := i3; i4 <= rune(57); i4++ {
						for i5 := i4; i5 <= rune(57); i5++ {
							p := []rune{i0, i1, i2, i3, i4, i5}
							double := double(p)
							inlim := inLimits(p)
							nodec := noDecreasing(p)
							if double && inlim && nodec {
								count++
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("count: ", count)
}

// 0 -> ascii decimal 48
// ...
// 9 -> ascii decimal 57

// Your puzzle input is 183564-657474
// Find how many passwords meet the criteria:
// 1. It is a six-digit number.
// 2. The value is within the range given in your puzzle input.
// 3. Two adjacent digits are the same (like 22 in 122345).
// 4. Going from left to right, the digits never decrease; they only ever
// increase or stay the same (like 111123 or 135679).
//
// Other than the range rule, the following are true:
// 111111 meets these criteria (double 11, never decreases).
// 223450 does not meet these criteria (decreasing pair of digits 50).
// 123789 does not meet these criteria (no double).
//
// How many different passwords within the range given in your puzzle input
// meet these criteria?
