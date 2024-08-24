package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
)

func rgb(c int) (int, int, int) {
	// original implementation from https://github.com/busyloop/lolcat/blob/master/lib/lolcat/lol.rb#L36
	freq := 0.1
	r := int(math.Sin(freq*float64(c)+0)*127 + 128)
	g := int(math.Sin(freq*float64(c)+2*math.Pi/3)*127 + 128)
	b := int(math.Sin(freq*float64(c)+4*math.Pi/3)*127 + 128)
	return r, g, b
}

func print(output []rune) {
	for j := 0; j < len(output); j++ {
		r, g, b := rgb(j)
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}

	fmt.Println()
}

func main() {
	info, _ := os.Stdin.Stat()
	var output []rune

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: echo \"Hello, world!\" | lolcat")
		os.Exit(1)
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
