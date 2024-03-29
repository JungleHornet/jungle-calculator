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
	"github.com/junglehornet/junglemath"
	"log"
	"os"
	"sort"
	"strconv"
)

func printHelp() {
	fmt.Println("\033[1;32mUsage: jcalc [command] [args]\033[0m")
	fmt.Println("\033[1;32mCommands:\033[0m")
	fmt.Println("\033[1;31mWARNING. THIS IS A DEV BUILD. SOME COMMANDS MAY NOT BE IMPLEMENTED.\033[0m")
	fmt.Println("    \033[1;32mjcalc -f [function] [args] - Use standalone functions \033[0m\n        pythag [leg 1 length] [leg 2 length] - Pythagorean Theorem Calculator \n        calc - Open general the Calculator")
	fmt.Println("    \033[1;32mjcalc -set [type] [name] [values] - Create a new variable or set the values of an existing one \033[0m\n        point [x] [y] [z (optional, default 0)] - Create a new point \n        line [point1] [point2] - Create a new line \n        triangle [point1] [point2] [point3] - Create a new triangle \n        angle [point1] [point2] [point3] - Create a new angle")
	fmt.Println("    \033[1;32mjcalc -vars [command] [args] - View/modify stored variables \033[0m\n        [no command] - View all variables and their values \n        clear - Clear all variables \n        delete [variable] - Delete a variable")
	fmt.Println("    \033[1;32mjcalc [variable] [function] - Do operations on a variable. [function] values: \033[0m\n        [no function] - View a variable and it's values \n        delete - Delete a variable \n        \033[1;32mVariable type specific:\033[0m")
	fmt.Println("            \033[1;32mLine:\033[0m \n                len/length - Measure the length of the line")
	fmt.Println("            \033[1;32mAngle:\033[0m \n                measure - Get the measure of the angle")
	fmt.Println("            \033[1;32mTriangle:\033[0m \n                orthocenter - Get the orthocenter of the triangle \n                circumcenter - Get the circumcenter of the triangle \n                centroid - Get the centroid of the triangle \n                incenter - Get the incenter of the triangle \n                orthocenter - Get the orthocenter of the triangle \n                parts - Get the info on each angle and side of the triangle.")
	fmt.Println("    \033[1;32mjcalc -help - Usage help")
}

func invCom(errCode int64) {
	if errCode == 0 {
		fmt.Println("\u001B[1;31mError: Invalid command. Run jcalc -help for usage.\033[0m")
	} else if errCode == 1 {
		fmt.Println("\033[1;31mError: Invalid variable. To view all variables, run jcalc -vars\033[0m")
	}
}

func main() {
	varfile, err := getVarfile()
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
							writeVar(varName, junglemath.Point{X: x, Y: y, Z: z}, varfile)
							fmt.Println("\033[1;32mSuccessfully set point " + varName + ".\033[0m")
							return
						} else if argLen > 5 {
							x, _ := strconv.ParseFloat(args[4], 64)
							y, _ := strconv.ParseFloat(args[5], 64)
							writeVar(varName, junglemath.Point{X: x, Y: y, Z: 0}, varfile)
							fmt.Println("\033[1;32mSuccessfully set point " + varName + ".\033[0m")
							return
						}
					case "line":
						if argLen > 5 {
							p1 := junglemath.ToPoint(getVarOfType(args[4], "junglemath.Point", varfile), args[4])
							p2 := junglemath.ToPoint(getVarOfType(args[5], "junglemath.Point", varfile), args[5])
							if getVarRaw(args[4], varfile) != nil && getVarRaw(args[5], varfile) != nil {
								writeVar(varName, junglemath.Line{P1: p1, P2: p2}, varfile)
								fmt.Println("\033[1;32mSuccessfully set line " + varName + ".\033[0m")
							} else {
								invCom(1)
							}
							return
						}
					case "triangle":
						if argLen > 6 {
							a := junglemath.ToPoint(getVarOfType(args[4], "junglemath.Point", varfile), args[4])
							b := junglemath.ToPoint(getVarOfType(args[5], "junglemath.Point", varfile), args[5])
							c := junglemath.ToPoint(getVarOfType(args[6], "junglemath.Point", varfile), args[6])
							if getVarRaw(args[4], varfile) != nil && getVarRaw(args[5], varfile) != nil && getVarRaw(args[6], varfile) != nil {
								if junglemath.IsValidTriangle(junglemath.Triangle{A: a, B: b, C: c}) {
									writeVar(varName, junglemath.Triangle{A: a, B: b, C: c}, varfile)
									fmt.Println("\033[1;32mSuccessfully set triangle " + varName + ".\033[0m")
								} else {
									fmt.Println("\033[1;31mError: Triangle " + args[4] + ", " + args[5] + ", " + args[6] + " is not a geometrically valid triangle. This may be because the angles do not add up to 180º or because of the triangle inequality theorem.\033[0m")
								}
							} else {
								invCom(1)
							}
							return
						}
					case "angle":
						if argLen > 6 {
							p1 := junglemath.ToPoint(getVarOfType(args[4], "junglemath.Point", varfile), args[4])
							p2 := junglemath.ToPoint(getVarOfType(args[5], "junglemath.Point", varfile), args[5])
							p3 := junglemath.ToPoint(getVarOfType(args[6], "junglemath.Point", varfile), args[6])
							if getVarRaw(args[4], varfile) != nil && getVarRaw(args[5], varfile) != nil && getVarRaw(args[6], varfile) != nil {
								writeVar(varName, junglemath.Angle{A: p1, B: p2, C: p3}, varfile)
								fmt.Println("\033[1;32mSuccessfully set angle " + varName + ".\033[0m")
							} else {
								invCom(1)
							}
							return
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
					return
				}
			} else {
				fmt.Println(string(varfile))
				return
			}
		case "-help":
			printHelp()
			return
		default:
			if getVarRaw(args[1], varfile) != nil {
				varMap := getVarRaw(args[1], varfile)
				varType := getVarType(varMap, args[1])
				if argLen > 2 {
					switch varType {
					case "junglemath.Line":
						lineVar := junglemath.ToLine(getVarOfType(args[1], varType, varfile), args[1])
						if args[2] == "len" || args[2] == "length" {
							fmt.Println("\033[1;32m", lineVar.Length(), "\033[0m")
							return
						}
					case "junglemath.Angle":
						angleVar := junglemath.ToAngle(getVarOfType(args[1], varType, varfile), args[1])
						if args[2] == "measure" {
							fmt.Println("\033[1;32m", angleVar.Measure(), "\033[0m")
							return
						}
					case "junglemath.Triangle":
						triangleVar := junglemath.ToTriangle(getVarOfType(args[1], varType, varfile), args[1])
						switch args[2] {
						case "orthocenter":
							fmt.Println("\033[1;32m", triangleVar.Orthocenter(), "\033[0m")
							return
						case "circumcenter":
							fmt.Println("\033[1;32m", triangleVar.Circumcenter(), "\033[0m")
							return
						case "centroid":
							fmt.Println("\033[1;32m", triangleVar.Centroid(), "\033[0m")
							return
						case "incenter":
							fmt.Println("\033[1;32m", triangleVar.Incenter(), "\033[0m")
							return
						case "parts":
							parts(triangleVar, args[1])
							return
						}
					}
					log.Fatal("Error: Feature not implemented yet. Make sure you have the latest build, or come back later once we add this feature.")
					return
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
				return
			}
		}
	}
	invCom(0)
}
