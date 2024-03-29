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
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/junglehornet/junglemath"
	"log"
	"os"
	"reflect"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func getVarfile() ([]byte, error) {
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

func getVarOfType(name string, typeString string, varfile []byte) map[string]any {
	var vars map[string]map[string]any
	err := json.Unmarshal(varfile, &vars)
	if err != nil {
		return nil
	}
	varMap := vars[name]
	storedVarTypeString := varMap["type"]
	delete(varMap, "type")
	_, success := storedVarTypeString.(string)
	if !success {
		log.Fatal("\033[1;31mError: Incorrect variable type: Variable " + name + " is of type " + storedVarTypeString.(string) + ", not of required type " + typeString + ".\033[0m")
	}
	if typeString == storedVarTypeString {
		return varMap
	} else {
		log.Fatal("\033[1;31mError: Invalid vars.json. Stored type of variable " + name + " is not a string. Please use \"jcalc -set\" to reset the fields of " + name + " or delete the invalid variable with \"jcalc " + name + " delete\"\033[0m")
	}
	return nil
}

func getVarRaw(name string, varfile []byte) map[string]any {
	var vars map[string]map[string]any
	err := json.Unmarshal(varfile, &vars)
	if err != nil {
		return nil
	}
	return vars[name]
}

func getVarType(m map[string]any, name string) string {
	varTypeString := m["type"]
	_, success := varTypeString.(string)
	if !success {
		log.Fatal("\033[1;31mError: Invalid vars.json. Stored type of variable " + name + " is not a string. Please use \"jcalc -set\" to reset the fields of " + name + " or delete the invalid variable with \"jcalc " + name + " delete\"\033[0m")
	}
	return varTypeString.(string)
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
		fmt.Println(err)
		return
	}
	vars[name] = varMap
	vars[name]["type"] = reflect.TypeOf(Var).String()
	marshaled, _ := json.Marshal(vars)
	homeDir, err := os.UserHomeDir()
	var indented bytes.Buffer
	err = json.Indent(&indented, marshaled, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(homeDir+"/jcalc/vars.json", indented.Bytes(), 0644)
	if err != nil {
		fmt.Println(err)

	}
}

func parts(t junglemath.Triangle, tName string) {
	fmt.Println("\033[1;32mParts of triangle " + tName + ":\033[0m")
	fmt.Println("    \033[1;32mSide lenths:\033[0m")
	sidea := junglemath.Line{P1: t.B, P2: t.C}
	fmt.Println("        a:", sidea.Length())
	sideb := junglemath.Line{P1: t.A, P2: t.C}
	fmt.Println("        b:", sideb.Length())
	sidec := junglemath.Line{P1: t.A, P2: t.B}
	fmt.Println("        c:", sidec.Length())
	fmt.Print("\n")
	fmt.Println("    \033[1;32mAngle measures:\033[0m")
	angleA := junglemath.Angle{A: t.B, B: t.A, C: t.C}
	fmt.Println("        A:", angleA.Measure())
	angleB := junglemath.Angle{A: t.A, B: t.B, C: t.C}
	fmt.Println("        angleB:", angleB.Measure())
	angleC := junglemath.Angle{A: t.A, B: t.C, C: t.B}
	fmt.Println("        angleC:", angleC.Measure())
	fmt.Print("\n")
	fmt.Println("    \033[1;32mSpecial points:\033[0m")
	fmt.Println("        Orthocenter:", t.Orthocenter())
	fmt.Println("        Circumcenter:", t.Circumcenter())
	fmt.Println("        Centroid:", t.Centroid())
	fmt.Println("        Incenter:", t.Incenter())
	fmt.Print("\n")
}
