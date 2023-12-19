package main

import (
	"fmt"
	"github.com/junglehornet/junglemath"
	"math"
	"strconv"
)

func simplifyRadical() bool {
	s := NewScanner()
	fmt.Println(d["str18"])
	root, err := strconv.ParseFloat(s.ReadLine(), 64)
	if root == 0 {
		fmt.Println(d["str4"])
		return true
	}
	if err != nil && root != 0 {
		fmt.Println(err)
		return true
	}

	simpleRoot := junglemath.SimplifyRadical(root)

	if math.Sqrt(root) == math.Trunc(math.Sqrt(root)) {
		simpleRoot = simpleRoot + d["str6"] + strconv.FormatFloat(math.Sqrt(root), 'f', -1, 64)
	}

	fmt.Println(simpleRoot)

	for {
		fmt.Println("\n" + d["str9"])
		yn := s.ReadLine()

		switch yn {
		case "y":
			fmt.Println(d["str12"])
			return true

		case "n":
			fmt.Println(d["n"] + d["str10"])
			return false

		default:
			fmt.Println(d["y"] + d["str6"] + d["n"] + d["str11"])
			return false
		}
	}

}
