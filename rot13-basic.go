package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getWordInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil { //error check
		return "-1"
	}
	input = strings.ToLower(strings.TrimSpace(input)) //remove whitespace and force lc
	for i := 0; i < len(input); i++ {
		if input[i] == 32 { //space is OK
			continue
		}
		if input[i] < 65 || input[i] > 122 || (input[i] > 90 && input[i] < 97) { //check for alpha
			return "-1"
		}
	}
	return input
}

func main() {

	reader := bufio.NewReader(os.Stdin) //reader setup

	fmt.Print("Enter text : ") //text to shift
	input := getWordInput(reader)
	for input == "-1" { //error check
		fmt.Print("Try again : ")
		input = getWordInput(reader)
	}

	var shifted string //final result
	for i := 0; i < len(input); i++ {
		if input[i] == 32 { //space is OK
			shifted += " "
			continue
		}
		shifted += (string)((int(input[i])-97+13)%26 + 97) //shift
	}
	fmt.Printf("ROT 13: %s\n", shifted) //result
}
