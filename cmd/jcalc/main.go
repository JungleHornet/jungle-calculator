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
	"encoding/json"
	"fmt"
	"github.com/junglehornet/junglemath"
	"os"
	"reflect"
	"strconv"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func getJson() ([]byte, error) {
	var homeDir string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	} else {
		if fileExists(homeDir + "/jcalc/vars.json") {
			return os.ReadFile(homeDir + "/jcalc/vars.json")
		} else {
			if fileExists(homeDir + "/jcalc/") {
				_, err = os.Create(homeDir + "/jcalc/vars.json")
			} else {
				err := os.Mkdir(homeDir+"/jcalc/", 0755)
				if err != nil {
					return nil, err
				}
				_, err = os.Create(homeDir + "/jcalc/vars.json")
			}
			if err != nil {
				return nil, err
			}
			return os.ReadFile(homeDir + "/jcalc/vars.json")
		}
	}
	return nil, nil
}

func getVar(name string, varfile []byte) any {
	var vars map[string]map[string]any
	err := json.Unmarshal(varfile, &vars)
	if err != nil {
		return nil
	}
	Var := vars[name]
	delete(Var, "type")
	return Var
}

func writeVar(name string, Var any, varfile []byte) {
	vars := make(map[string]map[string]any)
	err := json.Unmarshal(varfile, &vars)
	if err != nil {
		fmt.Println(err)
	}
	vars[name] = make(map[string]any)
	marshaledVar, _ := json.Marshal(Var)
	varMap := vars[name]
	err = json.Unmarshal(marshaledVar, &varMap)
	if err != nil {
		return
	}
	vars[name] = varMap
	vars[name]["type"] = reflect.TypeOf(Var).String()
	marshaled, _ := json.Marshal(vars)
	homeDir, err := os.UserHomeDir()
	err = os.WriteFile(homeDir+"/jcalc/vars.json", marshaled, 0644)
	if err != nil {
		fmt.Println(err)

	}
}

func invCom(errCode int64) {
	if errCode == 0 {
		fmt.Println("Error: Invalid command. Run jcalc -help for usage.")
	} else if errCode == 1 {
		fmt.Println("Error: Invalid variable. To view all variables, run jcalc -vars")
	}
}

func toPoint(m any) junglemath.Point {
	pointMap := m.(map[string]any)
	X := pointMap["X"].(float64)
	Y := pointMap["Y"].(float64)
	Z := pointMap["Z"].(float64)
	point := junglemath.Point{X: X, Y: Y, Z: Z}
	return point
}

func main() {
	varfile, err := getJson()
	// Use varFile so compiler shuts up
	_ = varfile
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
					fmt.Println("Calculator opened. Type q to exit.")
					junglemath.OpenCalculator()
					return
				case "pythag":
					if argLen > 4 {
						num1, _ := strconv.ParseFloat(args[3], 64)
						num2, _ := strconv.ParseFloat(args[4], 64)
						fmt.Println(junglemath.Pythag(num1, num2))
						return
					}
				}
			}
		case "-new":
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
							return
						} else if argLen > 5 {
							x, _ := strconv.ParseFloat(args[4], 64)
							y, _ := strconv.ParseFloat(args[5], 64)
							writeVar(varName, junglemath.Point{X: x, Y: y, Z: 0}, varfile)
							return
						}
					case "line":
						if argLen > 5 {
							p1 := toPoint(getVar(args[4], varfile))
							p2 := toPoint(getVar(args[5], varfile))
							if getVar(args[4], varfile) != nil && getVar(args[5], varfile) != nil {
								writeVar(varName, junglemath.Line{P1: p1, P2: p2}, varfile)
							} else {
								invCom(1)
							}
							return
						}
					case "triangle":
						if argLen > 6 {
							a := toPoint(getVar(args[4], varfile))
							b := toPoint(getVar(args[5], varfile))
							c := toPoint(getVar(args[6], varfile))
							if getVar(args[4], varfile) != nil && getVar(args[5], varfile) != nil && getVar(args[6], varfile) != nil {
								writeVar(varName, junglemath.Triangle{A: a, B: b, C: c}, varfile)
							} else {
								invCom(1)
							}
							return
						}
					case "angle":
						if argLen > 6 {
							p1 := toPoint(getVar(args[4], varfile))
							p2 := toPoint(getVar(args[5], varfile))
							p3 := toPoint(getVar(args[6], varfile))
							if getVar(args[4], varfile) != nil && getVar(args[5], varfile) != nil && getVar(args[6], varfile) != nil {
								writeVar(varName, junglemath.Angle{A: p1, B: p2, C: p3}, varfile)
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
			fmt.Println("Usage: jcalc [command] [args]")
			fmt.Println("Commands:")
			fmt.Println("WARNING. THIS IS A DEV BUILD. SOME COMMANDS MAY NOT BE IMPLEMENTED.")
			fmt.Println("    \033[1;32mjcalc -f [function] [args] - Use standalone functions \033[0m\n        pythag [leg 1 length] [leg 2 length] - Pythagorean Theorem Calculator \n        calc - Open general the Calculator")
			fmt.Println("    \033[1;32mjcalc -new [type] [name] [values] - Create a new variable \033[0m\n        point [x] [y] [z (optional, default 0)] - Create a new point \n        line [point1] [point2] - Create a new line \n        triangle [point1] [point2] [point3] - Create a new triangle \n        angle [point1] [point2] [point3] - Create a new angle")
			fmt.Println("    \033[1;32mjcalc -vars [command] [args] - View/modify stored variables \033[0m\n        [no command] - View all variables and their values \n        clear - Clear all variables \n        delete [variable] - Delete a variable")
			fmt.Println("    \033[1;32mjcalc [variable] [function] - Do operations on a variable \033[0m\n        [no function] - View a variable and it's values \n        set [values] - Set a variable's values \n        delete - Delete a variable \n        \033[1;32mVariable type specific:\033[0m")
			fmt.Println("            \u001B[1;32mLine:\u001B[0m \n                len/length - Measure the length of the line")
			fmt.Println("            \u001B[1;32mAngle:\u001B[0m \n                measure - Get the measure of the angle")
			fmt.Println("            \u001B[1;32mTriangle:\u001B[0m \n                orthocenter - Get the orthocenter of the triangle \n                circumcenter - Get the circumcenter of the triangle \n                centroid - Get the centroid of the triangle \n                incenter - Get the incenter of the triangle \n                orthocenter - Get the orthocenter of the triangle \n                parts - Get the info on each angle and side of the triangle.")
			fmt.Println("    \033[1;32mjcalc -help - Usage help")
			return
		}
	}
	invCom(0)
}
