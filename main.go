package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commands = map[string]func(string){
	"-c": countBytes,
	"-l": countLines,
	"-w": countWords,
	"-m": countCharacters,
}

func main() {
	// parse the command line arguments
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No arguments provided")
		os.Exit(1)
	}

	// if a command is provided, looks up the corresponding function and calls it
	if function, ok := commands[args[0]]; ok {
		function(args[1])
	} else {
		// if no command is provided, filename is taken as an argument, and all functions are called.
		countBytes(args[0])
		countLines(args[0])
		countWords(args[0])
	}
}

func countBytes(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Bytes: ", len(data))
}

func countLines(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	fmt.Println("Lines: ", lines)
}

func countWords(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	words := strings.Split(string(data), " ")
	fmt.Println("Words: ", len(words))
}

func countCharacters(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Characters: ", len([]rune(string(data))))
}
