package main

import (
	"fmt"
	"regexp"
)

func distance_calc() bool {
	fmt.Println(d["str2"])
	s := NewScanner()
	inpt := s.ReadLine()
	re := regexp.MustCompile(`(-?\\d*.?\\d*)\\s*,\\s*(-?\\d*.?\\d*)`)
	inpt1 := re.FindStringSubmatch(inpt)
	fmt.Println(len(inpt1))
	fmt.Println(inpt1)

	return true
}
