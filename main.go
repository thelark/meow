package main

import (
	"os"
	"strings"
	"io/ioutil"
	"fmt"
	"regexp"
)

func meow(input string) {
	input = strings.Replace(input, "喵", "", -1)
	lookup := map[string]string{
		".?": ">",
		"?.": "<",
		"..": "+",
		"!!": "-",
		"!.": ".",
		".!": ",",
		"!?": "[",
		"?!": "]",
	}
	reg, _ := regexp.Compile("[^\\.?!]+")
	input = string(reg.ReplaceAll([]byte(input), []byte("")))

	output := ""
	for i := 0; i < len(input); i += 2 {
		output += lookup[input[i:i+2]]
	}
	interpret(parse(output))
}

func main() {
	if len(os.Args) == 2 {
		c, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("read brainfuck file failed!\n")
			os.Exit(2)
		}
		meow(string(c))
	} else if len(os.Args) == 3 {
		if os.Args[1] == "-str" {
			str := os.Args[2]
			meow(str)
		} else {
			fmt.Printf("Usage: %s <brainfuck code>\n", os.Args[0])
			os.Exit(1)
		}
	} else {
		fmt.Printf("Usage: %s <brainfuck source file>\n", os.Args[0])
		os.Exit(1)
	}
}
