package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//// fmt.Println(string(f))
	l := strings.Split(string(f), "\n")

	if l[len(l)-1] == "" {
		l = l[:len(l)-1]
	}

	m := map[string][]string{}
	for _, k := range l {
		both := strings.Split(k, ")")
		_, ok := m[both[0]]
		if !ok {
			m[both[0]] = []string{both[1]}
		} else {
			exist := m[both[0]]
			exist = append(exist, both[1])
			m[both[0]] = exist
		}

	}

	orbits := 0
	find := "COM"
	allFound := []string{}
	found := []string{}
	distance := 0
	for {
		for {
			nf, ok := m[find]
			found = nf
			if !ok {
				if len(allFound) >= 2 {
					find = allFound[len(allFound)-2:][0]
				} else {
					find = allFound[0]
				}
				allFound = allFound[:len(allFound)-1]
				distance--
			} else {
				distance++
				break
			}
		}
		orbits += distance
		if len(found) == 1 {
			allFound = append(allFound, found[0])
			nf := found[0]
			delete(m, find)
			find = nf
		} else if len(found) > 1 {
			next := found[0]
			allFound = append(allFound, next)
			found = found[1:]
			m[find] = found
			find = next
		}
		if len(m) == 0 {
			break
		}
	}
	fmt.Println("orbits: ", orbits)
}
