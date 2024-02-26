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
