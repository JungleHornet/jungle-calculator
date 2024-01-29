package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/junglehornet/goScan"
	"github.com/junglehornet/junglemath"
)

func simplifyRadical() bool {
	s := goScan.NewScanner()
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

	return yn()

}
