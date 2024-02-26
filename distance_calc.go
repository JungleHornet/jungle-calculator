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
	"github.com/junglehornet/goScan"
	"github.com/junglehornet/junglemath"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func distanceCalc() bool {
	fmt.Println(d["str2"])
	s := goScan.NewScanner()
	inpt := s.ReadLine()
	re := regexp.MustCompile("\\(?(-?\\d*.?\\d*),(-?\\d*.?\\d*)\\)?")
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

	dist := junglemath.CalcDistance(x1, y1, x2, y2)

	rootStr, success := junglemath.CreateRoot(junglemath.CalcDistance(x1, y1, x2, y1))

	if !success {
		fmt.Println(d["str5"] + strconv.FormatFloat(dist, 'f', -1, 64))
	} else {

		rootStr, success := strings.CutPrefix(rootStr, "√")

		if !success {
			log.Fatal(d["str20"])
		}

		root, err := strconv.ParseFloat(rootStr, 64)

		if err != nil {
			fmt.Println(err)
		}

		radicand := root

		simpleRoot := junglemath.SimplifyRadical(radicand)

		simpleRootArray := strings.Split(simpleRoot, "√")

		simpleRootInt, _ := strconv.ParseFloat(simpleRootArray[0], 64)

		sqrtDist, _ := junglemath.CreateRoot(junglemath.CalcDistance(x1, y1, x2, y2))

		var response string
		if (math.Sqrt(root) == math.Trunc(math.Sqrt(root))) || simpleRootInt == root {
			response = d["str5"] + strconv.FormatFloat(dist, 'f', -1, 64) + d["str6"] + sqrtDist
		} else {
			response = d["str5"] + strconv.FormatFloat(dist, 'f', -1, 64) + d["str7"] + sqrtDist +
				d["str8"] + simpleRoot
		}
		fmt.Println(response)
	}

	return yn()
}

func distanceCalc3D() bool {
	fmt.Println(d["str21"])
	s := goScan.NewScanner()
	inpt := s.ReadLine()
	inpt = strings.ToLower(inpt)
	re := regexp.MustCompile(`\(?(-?\d*.?\d*),(-?\d*.?\d*),(-?\d*.?\d*)\)?`)
	inpt1 := re.FindStringSubmatch(inpt)
	if !(len(inpt1) == 4) {
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
	z1, err := strconv.ParseFloat(inpt1[3], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}

	fmt.Println(d["str22"])
	inpt = s.ReadLine()
	inpt2 := re.FindStringSubmatch(inpt)
	if !(len(inpt2) == 4) {
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
	z2, err := strconv.ParseFloat(inpt2[3], 64)
	if err != nil {
		fmt.Println(err)
		return true
	}

	dist := junglemath.CalcDistance3D(x1, y1, z1, x2, y2, z2)

	rootStr, success := junglemath.CreateRoot(junglemath.CalcDistance3D(x1, y1, z1, x2, y2, z2))

	if !success {
		fmt.Println(d["str5"] + strconv.FormatFloat(dist, 'f', -1, 64))
	} else {

		rootStr, success := strings.CutPrefix(rootStr, "√")

		if !success {
			log.Fatal(d["str20"])
		}

		root, err := strconv.ParseFloat(rootStr, 64)

		if err != nil {
			fmt.Println(err)
		}

		radicand := root

		simpleRoot := junglemath.SimplifyRadical(radicand)

		simpleRootArray := strings.Split(simpleRoot, "√")

		simpleRootInt, _ := strconv.ParseFloat(simpleRootArray[0], 64)

		sqrtDist, _ := junglemath.CreateRoot(junglemath.CalcDistance3D(x1, y1, z1, x2, y2, z2))

		var response string
		if (math.Sqrt(root) == math.Trunc(math.Sqrt(root))) || simpleRootInt == root {
			response = d["str5"] + strconv.FormatFloat(dist, 'f', -1, 64) + d["str6"] + sqrtDist
		} else {
			response = d["str5"] + strconv.FormatFloat(dist, 'f', -1, 64) + d["str7"] + sqrtDist +
				d["str8"] + simpleRoot
		}
		fmt.Println(response)
	}

	return yn()
}
