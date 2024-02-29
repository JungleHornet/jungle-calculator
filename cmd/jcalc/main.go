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
	vartype := Var["type"]
	return vartype
}

func writeVar(name string, Var any, varfile []byte) {
	vars := make(map[string]map[string]any)
	err := json.Unmarshal(varfile, &vars)
	if err != nil {
		fmt.Println(err)
	}
	vars[name] = make(map[string]any)
	vars[name]["type"] = reflect.TypeOf(Var).String()
	marshaled, err := json.Marshal(vars)
	if err != nil {
		return
	}
	homeDir, err := os.UserHomeDir()
	err = os.WriteFile(homeDir+"/jcalc/vars.json", marshaled, 0644)
	if err != nil {
		fmt.Println(err)

	}
}

func invCom() {
	fmt.Println("Error: Invalid command. Run jcalc -help for usage.")
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
		case "new":
			if argLen > 2 {
				varType := args[2]
				if argLen > 3 {
					varName := args[3]
					switch varType {
					case "point":
						if argLen > 5 {
							x, _ := strconv.ParseFloat(args[4], 64)
							y, _ := strconv.ParseFloat(args[5], 64)
							writeVar(varName, junglemath.Point{X: x, Y: y}, varfile)
							return
						}
						if argLen > 6 {
							x, _ := strconv.ParseFloat(args[4], 64)
							y, _ := strconv.ParseFloat(args[5], 64)
							z, _ := strconv.ParseFloat(args[6], 64)
							writeVar(varName, junglemath.Point{X: x, Y: y, Z: z}, varfile)
							return
						}
					case "line":
						if argLen > 5 {
							p1 := getVar(args[4], varfile)
							p2 := getVar(args[5], varfile)
							point1 := p1.(junglemath.Point)
							point2 := p2.(junglemath.Point)
							writeVar(varName, junglemath.Line{P1: point1, P2: point2}, varfile)
							return
						}
					case "triangle":
						return
					case "angle":
						return
					}
				}
			}
		}
	}
	invCom()
}
