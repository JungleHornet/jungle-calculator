package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
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

	x1, err := strconv.ParseFloat(inpt1[1], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}
	y1, err := strconv.ParseFloat(inpt1[2], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}

	fmt.Println(d["str3"])
	inpt = s.ReadLine()
	inpt2 := re.FindStringSubmatch(inpt)
	if !(len(inpt2) == 3) {
		fmt.Println(d["str4"])
		return true
	}

	x2, err := strconv.ParseFloat(inpt2[1], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}
	y2, err := strconv.ParseFloat(inpt2[2], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}

	a := x2 - x1
	b := y2 - y1

	a = a * a
	b = b * b

	dist := math.Sqrt(a + b)

	// This is commented so Golang will work because I haven't used sqrtDist yet
	// sqrtDist := "√" + strconv.FormatFloat(math.Round(dist*dist), 'f', -1, 64)
	root := math.Round(dist * dist)

	rootCoefficient := 1
	simpleRootInt := root

	for i := float64(2); i <= math.Round(math.Sqrt(root)); i++ {
		if (simpleRootInt / i) == math.Trunc(simpleRootInt/i) {
			if !(i == simpleRootInt) && !(i == 1) {
				if math.Sqrt(simpleRootInt/i) == math.Trunc(math.Sqrt(simpleRootInt/i)) {
					simpleRootInt = i
					rootCoefficient = rootCoefficient * int(math.Sqrt(root/i))
				}
			}
		}
	}

	return true
}
