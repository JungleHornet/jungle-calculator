package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func pythag() bool {
	s := NewScanner()
	fmt.Println(d["str15"])
	inpt := s.ReadLine()

	re := regexp.MustCompile("\\d*.?\\d*")
	inpt1 := re.FindString(inpt)
	if inpt1 == "" {
		fmt.Println(d["str4"])
		return true
	}

	leg1, err := strconv.ParseFloat(inpt1, 64)
	if err != nil {
		fmt.Println(err)
		return true
	}

	fmt.Println(d["str16"])
	inpt = s.ReadLine()
	inpt2 := re.FindString(inpt)
	if inpt2 == "" {
		fmt.Println(d["str4"])
		return true
	}

	leg2, err := strconv.ParseFloat(inpt2, 64)

	leg1 = leg1 * leg1
	leg2 = leg2 * leg2

	hyp := math.Sqrt(leg1 + leg2)

	root := math.Round(hyp * hyp)
	sqrtHyp := "√" + strconv.FormatFloat(root, 'f', -1, 64)

	rootCoefficient := int64(1)
	simpleRootInt := root

	for i := float64(2); i <= math.Round(math.Sqrt(root)); i++ {
		if (simpleRootInt / i) == math.Trunc(simpleRootInt/i) {
			if !(i == simpleRootInt) && !(i == 1) {
				if math.Sqrt(simpleRootInt/i) == math.Trunc(math.Sqrt(simpleRootInt/i)) {
					simpleRootInt = i
					rootCoefficient = rootCoefficient * int64(math.Sqrt(root/i))
				}
			}
		}
	}

	simpleRoot := strconv.FormatInt(rootCoefficient, 10) + "√" +
		strconv.FormatFloat(simpleRootInt, 'f', -1, 64)

	var response string
	if (math.Sqrt(root) == math.Trunc(math.Sqrt(root))) || simpleRootInt == root {
		response = d["str5"] + strconv.FormatFloat(hyp, 'f', -1, 64) + d["str6"] + sqrtHyp
	} else {
		response = d["str5"] + strconv.FormatFloat(hyp, 'f', -1, 64) + d["str7"] + sqrtHyp +
			d["str8"] + simpleRoot
	}

	fmt.Println(response)

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
