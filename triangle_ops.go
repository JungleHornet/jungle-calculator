/*
JungleCalculator - An open-source Go calculator for advanced math functions.
Copyright (c) 2023-present  JungleHornet

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/junglehornet/goScan"
	"github.com/junglehornet/junglemath"
)

func pythag() bool {
	s := goScan.NewScanner()
	fmt.Println(d["str15"])
	inpt := s.ReadLine()

	re := regexp.MustCompile(`\s*(-?\s*\d*\s*.?\s*\d*)\s*`)
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

		sqrtInt := root

		simpleSqrtHyp := junglemath.SimplifyRadical(sqrtInt)

		simpRootParts := strings.Split(simpleSqrtHyp, "√")

		simpleRootInt, err := strconv.ParseFloat(simpRootParts[1], 64)
		if err != nil {
			fmt.Println(err)
		}

		rootStr = "√" + rootStr

		var response string
		if (math.Sqrt(root) == math.Trunc(math.Sqrt(root))) || simpleRootInt == root {
			response = d["str5"] + strconv.FormatFloat(hyp, 'f', -1, 64) + d["str6"] + rootStr
		} else {
			response = d["str5"] + strconv.FormatFloat(hyp, 'f', -1, 64) + d["str7"] + rootStr +
				d["str8"] + simpleSqrtHyp
		}

		fmt.Println(response)
	}

	return yn()
}

func orthocenter() bool {
	fmt.Println(d["str2"])
	s := goScan.NewScanner()
	inpt := s.ReadLine()
	re := regexp.MustCompile(`\(?(-?\d*.?\d*),(-?\d*.?\d*)\)?`)
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

	fmt.Println(d["str24"])
	inpt = s.ReadLine()
	inpt3 := re.FindStringSubmatch(inpt)
	if !(len(inpt2) == 3) {
		fmt.Println(d["str4"])
		return true
	}

	x3, err := strconv.ParseFloat(inpt3[1], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}
	y3, err := strconv.ParseFloat(inpt3[2], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}

	x, y := junglemath.GetOrthocenter(x1, y1, x2, y2, x3, y3)

	fmt.Println(d["str25"] + "(" + strconv.FormatFloat(x, 'f', -1, 64) + ", " + strconv.FormatFloat(y, 'f', -1, 64) + ")")

	return yn()
}

func circumcenter() bool {
	fmt.Println(d["str2"])
	s := goScan.NewScanner()
	inpt := s.ReadLine()
	re := regexp.MustCompile(`\(?(-?\d*.?\d*),(-?\d*.?\d*)\)?`)
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

	fmt.Println(d["str24"])
	inpt = s.ReadLine()
	inpt3 := re.FindStringSubmatch(inpt)
	if !(len(inpt2) == 3) {
		fmt.Println(d["str4"])
		return true
	}

	x3, err := strconv.ParseFloat(inpt3[1], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}
	y3, err := strconv.ParseFloat(inpt3[2], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}

	x, y := junglemath.GetCircumcenter(x1, y1, x2, y2, x3, y3)

	fmt.Println(d["str26"] + "(" + strconv.FormatFloat(x, 'f', -1, 64) + ", " + strconv.FormatFloat(y, 'f', -1, 64) + ")")

	return yn()
}

func centroid() bool {
	fmt.Println(d["str2"])
	s := goScan.NewScanner()
	inpt := s.ReadLine()
	re := regexp.MustCompile(`\(?(-?\d*.?\d*),(-?\d*.?\d*)\)?`)
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

	fmt.Println(d["str24"])
	inpt = s.ReadLine()
	inpt3 := re.FindStringSubmatch(inpt)
	if !(len(inpt2) == 3) {
		fmt.Println(d["str4"])
		return true
	}

	x3, err := strconv.ParseFloat(inpt3[1], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}
	y3, err := strconv.ParseFloat(inpt3[2], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}

	x, y := junglemath.GetCentroid(x1, y1, x2, y2, x3, y3)

	fmt.Println(d["str27"] + "(" + strconv.FormatFloat(x, 'f', -1, 64) + ", " + strconv.FormatFloat(y, 'f', -1, 64) + ")")

	return yn()
}
func incenter() bool {
	fmt.Println(d["str2"])
	s := goScan.NewScanner()
	inpt := s.ReadLine()
	re := regexp.MustCompile(`\(?(-?\d*.?\d*),(-?\d*.?\d*)\)?`)
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

	fmt.Println(d["str24"])
	inpt = s.ReadLine()
	inpt3 := re.FindStringSubmatch(inpt)
	if !(len(inpt2) == 3) {
		fmt.Println(d["str4"])
		return true
	}

	x3, err := strconv.ParseFloat(inpt3[1], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}
	y3, err := strconv.ParseFloat(inpt3[2], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}

	x, y := junglemath.GetIncenter(x1, y1, x2, y2, x3, y3)

	fmt.Println(d["str28"] + "(" + strconv.FormatFloat(x, 'f', -1, 64) + ", " + strconv.FormatFloat(y, 'f', -1, 64) + ")")

	return yn()
}
