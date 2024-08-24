package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func print(output []rune) {
	colors := []string{}

	for i := 31; i < 37; i++ {
		s := "\033[" + fmt.Sprintf("%d", i) + "m"
		colors = append(colors, s)
	}

	for j := 0; j < len(output); j++ {
		fmt.Printf("%s%c", colors[j%len(colors)], output[j])
	}

	fmt.Println()
}

func main() {
	info, _ := os.Stdin.Stat()
	var output []rune

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo \"Hello, world!\" | lolcat")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	print(output)
}
