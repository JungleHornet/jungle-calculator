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
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/junglehornet/goScan"
	"github.com/junglehornet/junglemath"
	"os"
	"strconv"
	"strings"
)

//go:embed langs
var f embed.FS
var en, hu []byte

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func init() {
	en, _ = f.ReadFile("langs/en.json")
	hu, _ = f.ReadFile("langs/hu.json")
}

var d map[string]string

func run(myFunc func() bool) {
	for myFunc() {
	}
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
			_, err := os.Create(homeDir + "/jcalc/vars.json")
			if err != nil {
				return nil, err
			}
			return os.ReadFile(homeDir + "/jcalc/vars.json")
		}
	}
	return nil, nil
}

/*
func getStoredVar(name string, varfile []byte) StoredVar {
	m := map[string]StoredVar{}
	err := json.Unmarshal(varfile, &m)
	if err != nil {
		return StoredVar{}
	}
	for i := range m {
		if m[i].Name == name {
			return m[i]
		}
	}
	return StoredVar{}
}
*/

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

func main_loop() {
	for {
		fmt.Println(d["str13"])
		fmt.Println(d["func1"])
		fmt.Println(d["func2"])
		fmt.Println(d["func3"])
		fmt.Println(d["func4"])
		fmt.Println(d["quit"])
		s := goScan.NewScanner()
		inpt := strings.ToLower(s.ReadLine())

		switch inpt {
		case "1":
			fmt.Print("\n\n")
			fmt.Println(d["str23"])
			junglemath.OpenCalculator()
		case "2":
			fmt.Print("\n\n")
			func2Loop(s)
		case "3":
			fmt.Print("\n\n")
			func3Loop(s)
		case "4":
			fmt.Print("\n\n")
			func4Loop(s)
		case "q":
			fmt.Println(d["quit"] + d["str10"])
			os.Exit(0)
		default:
			fmt.Println(d["str14"])

		}
	}
}

func func2Loop(s *goScan.Scanner) {
	for {
		fmt.Println(d["str13"])
		fmt.Println(d["dist1"])
		fmt.Println(d["dist2"])
		fmt.Println(d["back"])
		inpt := strings.ToLower(s.ReadLine())
		switch inpt {
		case "1":
			fmt.Print("\n\n")
			run(distanceCalc)
		case "2":
			fmt.Print("\n\n")
			run(distanceCalc3D)
		case "b":
			fmt.Print("\n\n")
			return
		}
	}
}

func func3Loop(s *goScan.Scanner) {
	for {
		fmt.Println(d["str13"])
		fmt.Println(d["tri1"])
		fmt.Println(d["tri2"])
		fmt.Println(d["tri3"])
		fmt.Println(d["tri4"])
		fmt.Println(d["tri5"])
		fmt.Println(d["back"])
		inpt := strings.ToLower(s.ReadLine())
		switch inpt {
		case "1":
			fmt.Print("\n\n")
			run(pythag)
		case "2":
			fmt.Print("\n\n")
			run(orthocenter)
		case "3":
			fmt.Print("\n\n")
			run(circumcenter)
		case "4":
			fmt.Print("\n\n")
			run(centroid)
		case "5":
			fmt.Print("\n\n")
			run(incenter)
		case "b":
			fmt.Print("\n\n")
			return
		}
	}
}

func func4Loop(s *goScan.Scanner) {
	for {
		fmt.Println(d["str13"])
		fmt.Println(d["misc1"])
		fmt.Println(d["back"])
		inpt := strings.ToLower(s.ReadLine())
		switch inpt {
		case "1":
			fmt.Print("\n\n")
			run(simplifyRadical)
		case "b":
			fmt.Print("\n\n")
			return
		}
	}
}

func yn() bool {
	s := goScan.NewScanner()
	fmt.Println("\n" + d["str9"])
	yn := s.ReadLine()

	switch yn {
	case "y":
		fmt.Println(d["str12"])
		return true

	case "n":
		fmt.Println(d["n"] + d["str10"])
		return false

	default:
		fmt.Println(d["y"] + d["str6"] + d["n"] + d["str11"])
		return false
	}
}
