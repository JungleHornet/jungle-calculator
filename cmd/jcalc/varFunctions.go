package main

import "C"
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
		log.Fatal("Error: Incorrect variable type: Variable " + name + " is of type " + storedVarTypeString.(string) + ", not of required type " + typeString + ".")
	}
	if typeString == storedVarTypeString {
		return varMap
	} else {
		log.Fatal("Error: Invalid vars.json. Stored type of variable " + name + " is not a string. Please use \"jcalc -set\" to reset the fields of " + name + " or delete the invalid variable with \"jcalc " + name + " delete\"")
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
		log.Fatal("Error: Invalid vars.json. Stored type of variable " + name + " is not a string. Please use \"jcalc -set\" to reset the fields of " + name + " or delete the invalid variable with \"jcalc " + name + " delete\"")
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

func isValidPoint(m map[string]any) bool {
	delete(m, "type")
	if reflect.TypeOf(m["X"]).Kind() != reflect.Float64 {
		return false
	}
	if reflect.TypeOf(m["Y"]).Kind() != reflect.Float64 {
		return false
	}
	if reflect.TypeOf(m["Z"]).Kind() != reflect.Float64 {
		return false
	}
	return true
}

func toPoint(m map[string]any, name string) junglemath.Point {
	if isValidPoint(m) {
		X := m["X"].(float64)
		Y := m["Y"].(float64)
		Z := m["Z"].(float64)
		point := junglemath.Point{X: X, Y: Y, Z: Z}
		return point
	}
	log.Fatal("Error: Invalid stored variable: " + name + ". Please use jcalc -set to reset the variable's values.")
	return junglemath.Point{}
}

func isValidLine(m map[string]any) bool {
	P1, success := m["P1"].(map[string]any)
	if !success {
		return false
	}
	if !isValidPoint(P1) {
		return false
	}
	P2, success := m["P2"].(map[string]any)
	if !success {
		return false
	}
	if !isValidPoint(P2) {
		return false
	}
	return true
}

func toLine(m map[string]any, name string) junglemath.Line {
	if isValidLine(m) {
		P1 := toPoint(m["P1"].(map[string]any), name)
		P2 := toPoint(m["P2"].(map[string]any), name)
		line := junglemath.Line{P1: P1, P2: P2}
		return line
	}
	log.Fatal("Error: Invalid stored variable: " + name + ". Please use jcalc -set to reset the variable's values.")
	return junglemath.Line{}
}

func isValidAngle(m map[string]any) bool {
	A, success := m["A"].(map[string]any)
	if !success {
		return false
	}
	if !isValidPoint(A) {
		return false
	}
	B, success := m["B"].(map[string]any)
	if !success {
		return false
	}
	if !isValidPoint(B) {
		return false
	}
	C, success := m["C"].(map[string]any)
	if !success {
		return false
	}
	if !isValidPoint(C) {
		return false
	}
	return true
}

func toAngle(m map[string]any, name string) junglemath.Angle {
	if isValidAngle(m) {
		A := toPoint(m["A"].(map[string]any), name)
		B := toPoint(m["B"].(map[string]any), name)
		C := toPoint(m["C"].(map[string]any), name)
		angle := junglemath.Angle{A: A, B: B, C: C}
		return angle
	}
	log.Fatal("Error: Invalid stored variable: " + name + ". Please use jcalc -set to reset the variable's values.")
	return junglemath.Angle{}
}

func isValidTriangle(m map[string]any) bool {
	A, success := m["A"].(map[string]any)
	if !success {
		return false
	}
	if !isValidPoint(A) {
		return false
	}
	B, success := m["B"].(map[string]any)
	if !success {
		return false
	}
	if !isValidPoint(B) {
		return false
	}
	C, success := m["C"].(map[string]any)
	if !success {
		return false
	}
	if !isValidPoint(C) {
		return false
	}
	return true
}

func toTriangle(m map[string]any, name string) junglemath.Triangle {
	if isValidAngle(m) {
		A := toPoint(m["A"].(map[string]any), name)
		B := toPoint(m["B"].(map[string]any), name)
		C := toPoint(m["C"].(map[string]any), name)
		angle := junglemath.Triangle{A: A, B: B, C: C}
		return angle
	}
	log.Fatal("Error: Invalid stored variable: " + name + ". Please use jcalc -set to reset the variable's values.")
	return junglemath.Triangle{}
}
