package main

import (
	"fmt"
	"regexp"
)

func distanceCalc() bool {
	fmt.Println(d["str2"])
	s := NewScanner()
	inpt := s.ReadLine()
	re := regexp.MustCompile("(-?\\d*.?\\d*)\\s*,\\s*(-?\\d*.?\\d*)")
	inpt1 := re.FindStringSubmatch(inpt)
	if !(len(inpt1) == 3) {
		fmt.Println(d["str4"])
		return true
	}

	return true
}
