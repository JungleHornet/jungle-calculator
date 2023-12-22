package main

import (
	"fmt"
	"github.com/junglehornet/junglemath"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func pythag() bool {
	s := NewScanner()
	fmt.Println(d["str15"])
	inpt := s.ReadLine()

	re := regexp.MustCompile("\\s*(-?\\s*\\d*\\s*.?\\s*\\d*)\\s*")
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

	if err != nil {
		fmt.Println(err)
	}

	hyp := junglemath.Pythag(leg1, leg2)

	rootStr, success := junglemath.CreateRoot(junglemath.Pythag(leg1, leg2))

	if !success {
		fmt.Println(d["str5"] + strconv.FormatFloat(hyp, 'f', -1, 64))
	} else {

		rootStr, success := strings.CutPrefix(rootStr, "√")

		if !success {
			log.Fatal(d["str20"] + "\"" + rootStr + "\"")
		}

		root, err := strconv.ParseFloat(rootStr, 64)

		if err != nil {
			fmt.Println(err)
		}

		simpleSqrtHyp, _ := junglemath.CreateRoot(junglemath.Pythag(leg1, leg2))

		sqrtInt := math.Sqrt(root)

		simpleSqrtHyp = junglemath.SimplifyRadical(sqrtInt)

		simpRootParts := strings.Split(simpleSqrtHyp, "√")

		simpleRootInt, err := strconv.ParseFloat(simpRootParts[1], 64)
		if err != nil {
			fmt.Println(err)
		}

		var response string
		if (math.Sqrt(root) == math.Trunc(math.Sqrt(root))) || simpleRootInt == root {
			response = d["str5"] + strconv.FormatFloat(hyp, 'f', -1, 64) + d["str6"] + rootStr
		} else {
			response = d["str5"] + strconv.FormatFloat(hyp, 'f', -1, 64) + d["str7"] + rootStr +
				d["str8"] + simpleSqrtHyp
		}

		fmt.Println(response)
	}

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
