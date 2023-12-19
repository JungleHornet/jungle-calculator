package main

import (
	"bufio"
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

//go:embed langs
var f embed.FS
var en, hu []byte

func init() {
	en, _ = f.ReadFile("langs/en.json")
	hu, _ = f.ReadFile("langs/hu.json")
}

var d map[string]string

func run(myFunc func() bool) {
	for myFunc() {
	}
}

type Scanner struct {
	scanner *bufio.Scanner
}

func NewScanner() *Scanner {
	m := &Scanner{
		scanner: bufio.NewScanner(os.Stdin),
	}
	return m
}

func (m *Scanner) ReadLine() string {
	m.scanner.Scan()
	err := m.scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return m.scanner.Text()
}

func main() {
	fmt.Println("Please select a language (english (en), magyar (ma) )")

	s := NewScanner()

	langInpt := strings.ToLower(s.ReadLine())
	var dictFile []byte
	switch langInpt {
	case "en":
		fmt.Println("English selected.")
		dictFile = en

	case "ma":
		fmt.Println("Magyar válogatott.")
		dictFile = hu

	default:
		fmt.Println("Language not recognised, defaulting to english.")
		dictFile = en
	}

	err := json.Unmarshal(dictFile, &d)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(d["str1"])
	main_loop()
}

func main_loop() {
	for {
		fmt.Println(d["str13"])
		fmt.Println(d["func1"])
		fmt.Println(d["func2"])
		fmt.Println(d["func3"])
		fmt.Println(d["quit"])
		s := NewScanner()
		inpt := strings.ToLower(s.ReadLine())

		switch inpt {
		case "1":
			fmt.Println("\n\n")
			run(distanceCalc)
		case "2":
			fmt.Println("\n\n")
			run(distanceCalc3D)
		case "3":
			fmt.Println("\n\n")
			run(pythag)
		case "4":
			fmt.Println("\n\n")
			run(simplifyRadical)
		case "q":
			fmt.Println(d["quit"] + d["str10"])
			os.Exit(0)
		default:
			fmt.Println(d["str14"])

		}
	}
}
