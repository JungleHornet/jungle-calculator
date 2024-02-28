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
	"strconv"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func getJson() ([]byte, error) {
	var homeDir string
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	} else {
		if fileExists(homeDir + "/jcalc/vars.json") {
			fmt.Println("file exists")
			return os.ReadFile(homeDir + "/jcalc/vars.json")
		} else {
			fmt.Println("file does not exist 1")
			if fileExists(homeDir + "/jcalc/") {
				_, err = os.Create(homeDir + "/jcalc/vars.json")
			} else {
				os.Mkdir(homeDir+"/jcalc/", 0755)
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

func getVar(name string, varfile []byte) interface{} {
	var vars map[string]map[string]any
	err := json.Unmarshal(varfile, &vars)
	if err != nil {
		return nil
	}
	Var := vars[name]
	vartype := Var["type"]
	return vartype
}

func main() {
	varfile, err := getJson()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(getVar("test", varfile))
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "-f":
			if len(args) > 2 {
				switch args[2] {
				case "calc":
					junglemath.OpenCalculator()
					return
				case "pythag":
					if len(args) > 4 {
						num1, _ := strconv.ParseFloat(args[3], 64)
						num2, _ := strconv.ParseFloat(args[4], 64)
						fmt.Println(junglemath.Pythag(num1, num2))
						return
					}
				}
			}
		}
	}
}
