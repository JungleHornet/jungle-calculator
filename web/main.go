/*
JungleCalculatorWASM - An open-source Go-WASM calculator for advanced math functions.
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

package jcalc_wasm

import (
	"fmt"
	"github.com/junglehornet/junglemath"
	"log"
	"os"
	"sort"
	"strconv"
)

type jcalcRes struct {
	varFile []byte
	runError bool
	resMessage string
	result string
}

func Run(varFile []byte, opId int, opArgs map[string]any) jcalcRes  {
	var res jcalcRes
	if err != nil {
		fmt.Println(err)
	}
	args := os.Args
	argLen := len(args)
	if argLen > 1 {
		switch args[1] {
		case "-f":
			if argLen > 2 {
				switch args[2] {
				case "calc":
					fmt.Println("\033[1;32mCalculator opened. Type q to exit\033[0m")
					junglemath.OpenCalculator()
					return
				case "pythag":
					if argLen > 4 {
						num1, _ := strconv.ParseFloat(args[3], 64)
						num2, _ := strconv.ParseFloat(args[4], 64)
						fmt.Println("\033[1;32m", junglemath.Pythag(num1, num2), "\033[0m")
						return
					}
				}
			}
		case "-set":
			if argLen > 2 {
				varType := args[2]
				if argLen > 3 {
					varName := args[3]
					switch varType {
					case "point":
						if argLen > 6 {
							x, _ := strconv.ParseFloat(args[4], 64)
							y, _ := strconv.ParseFloat(args[5], 64)
							z, _ := strconv.ParseFloat(args[6], 64)
							res.varFile, err = writeVar(varName, junglemath.Point{X: x, Y: y, Z: z}, varFile)
							if err {
							res.runError = true
							res.resMessage = "Error encountered. Please open an issue at https://github.com/JungleHornet/jungle-calculator/issues/new if this keeps happening."
							}
						} else if argLen > 5 {
							x, _ := strconv.ParseFloat(args[4], 64)
							y, _ := strconv.ParseFloat(args[5], 64)
							res.varFile, err = writeVar(varName, junglemath.Point{X: x, Y: y, Z: 0}, varFile)
							if err {
							res.runError = true
							res.resMessage = "Error encountered. Please open an issue at https://github.com/JungleHornet/jungle-calculator/issues/new if this keeps happening."
						}
					}
					case "line":
						if argLen > 5 {
							p1 := junglemath.ToPoint(getVarOfType(args[4], "junglemath.Point", varFile), args[4])
							p2 := junglemath.ToPoint(getVarOfType(args[5], "junglemath.Point", varFile), args[5])
							if getVarRaw(args[4], varFile) != nil && getVarRaw(args[5], varFile) != nil {
								res.varFile, err = writeVar(varName, junglemath.Point{X: x, Y: y, Z: 0}, varfile)
								if err {
								res.runError = true
								res.resMessage = "Error encountered. Please open an issue at https://github.com/JungleHornet/jungle-calculator/issues/new if this keeps happening."
								}
							} else {
								res.runError = true
								res.resMessage = "Error: Invalid variable. To view all variables, run jcalc -vars"
							}
						}
					case "triangle":
						if argLen > 6 {
							a := junglemath.ToPoint(getVarOfType(args[4], "junglemath.Point", varFile), args[4])
							b := junglemath.ToPoint(getVarOfType(args[5], "junglemath.Point", varFile), args[5])
							c := junglemath.ToPoint(getVarOfType(args[6], "junglemath.Point", varFile), args[6])
							if getVarRaw(args[4], varFile) != nil && getVarRaw(args[5], varFile) != nil && getVarRaw(args[6], varFile) != nil {
								if junglemath.IsValidTriangle(junglemath.Triangle{A: a, B: b, C: c}) {
									writeVar(varName, junglemath.Triangle{A: a, B: b, C: c}, varFile)
									fmt.Println("\033[1;32mSuccessfully set triangle " + varName + ".\033[0m")
								} else {
									fmt.Println("\033[1;31mError: Triangle " + args[4] + ", " + args[5] + ", " + args[6] + " is not a geometrically valid triangle. This may be because the angles do not add up to 180º or because of the triangle inequality theorem.\033[0m")
								}
							} else {
								invCom(1)
							}
						}
					case "angle":
						if argLen > 6 {
							p1 := junglemath.ToPoint(getVarOfType(args[4], "junglemath.Point", varFile), args[4])
							p2 := junglemath.ToPoint(getVarOfType(args[5], "junglemath.Point", varFile), args[5])
							p3 := junglemath.ToPoint(getVarOfType(args[6], "junglemath.Point", varFile), args[6])
							if getVarRaw(args[4], varFile) != nil && getVarRaw(args[5], varFile) != nil && getVarRaw(args[6], varFile) != nil {
								writeVar(varName, junglemath.Angle{A: p1, B: p2, C: p3}, varFile)
								fmt.Println("\033[1;32mSuccessfully set angle " + varName + ".\033[0m")
							}
						}
					}
				}
			}
		case "-vars":
			if argLen > 2 {
				if args[2] == "clear" {
					homeDir, err := os.UserHomeDir()
					err = os.WriteFile(homeDir+"/jcalc/vars.json", []byte(""), 0644)
					if err != nil {
						fmt.Println(err)
					}
				}
			} else {
				fmt.Println(string(varFile))
			}
		default:
			if getVarRaw(args[1], varFile) != nil {
				varMap := getVarRaw(args[1], varFile)
				varType := getVarType(varMap, args[1])
				if argLen > 2 {
					switch varType {
					case "junglemath.Line":
						lineVar := junglemath.ToLine(getVarOfType(args[1], varType, varFile), args[1])
						if args[2] == "len" || args[2] == "length" {
							fmt.Println("\033[1;32m", lineVar.Length(), "\033[0m")
						}
					case "junglemath.Angle":
						angleVar := junglemath.ToAngle(getVarOfType(args[1], varType, varFile), args[1])
						if args[2] == "measure" {
							fmt.Println("\033[1;32m", angleVar.Measure(), "\033[0m")
						}
					case "junglemath.Triangle":
						triangleVar := junglemath.ToTriangle(getVarOfType(args[1], varType, varFile), args[1])
						switch args[2] {
						case "orthocenter":
							fmt.Println("\033[1;32m", triangleVar.Orthocenter(), "\033[0m")
						case "circumcenter":
							fmt.Println("\033[1;32m", triangleVar.Circumcenter(), "\033[0m")
						case "centroid":
							fmt.Println("\033[1;32m", triangleVar.Centroid(), "\033[0m")
						case "incenter":
							fmt.Println("\033[1;32m", triangleVar.Incenter(), "\033[0m")
						case "parts":
							parts(triangleVar, args[1])
						}
					}
					log.Fatal("Error: Feature not implemented yet. Make sure you have the latest build, or come back later once we add this feature.")
				}
				fmt.Println("Variable " + args[1] + ":")
				fmt.Println("Type:", varType)
				delete(varMap, "type")
				var names []string
				for name := range varMap {
					names = append(names, name)
				}
				sort.Slice(names, func(i, j int) bool { return names[i] < names[j] })
				for _, name := range names {
					fmt.Println(name+":", varMap[name])
				}
			}
		}
	}

	return res
}
