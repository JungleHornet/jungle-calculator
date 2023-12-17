package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

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
	var dictFile string
	switch langInpt {
	case "en":
		fmt.Println("English selected.")
		dictFile = "en.json"

	case "ma":
		fmt.Println("Magyar válogatott.")
		dictFile = "ma.json"

	default:
		fmt.Println("Language not recognised, defaulting to english.")
		dictFile = "en.json"
	}

	dictFile = "langs/" + dictFile

	jsonFile, err := os.Open(dictFile)
	if err != nil {
		fmt.Println(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)

	all, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(all, &d)
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

		case "2":
			fmt.Println("\n\n")

		case "3":
			fmt.Println("\n\n")

		case "q":
			fmt.Println(d["quit"] + d["str10"])
			os.Exit(0)

		default:
			fmt.Println(d["str14"])

		}
	}
}
