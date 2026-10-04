package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var cmdLoop bool = true
	for cmdLoop {
		scanner.Scan()
		cmdLoop = processCmd(scanner.Text())
	}
}

// err prints an error message.
func err(text string) {
	fmt.Print("error: " + text + "\n")
}

// processCmd parses input and calls the respective function with arguments
func processCmd(cmdString string) bool {
	cmdItems := strings.Fields(cmdString)
	switch cmdItems[0] {
	case "base", "b":
		cmdBase(cmdItems[1:])
	case "2c":
		cmd2c(cmdItems[1:])
	case "2cd":
		cmd2cd(cmdItems[1:])
	case "exit":
		return false
	default:
	}
	return true
}

// cmdBase calls the correct base conversion function based on the input
func cmdBase(items []string) {
	// Items: sourceBase, destBase, value
	if len(items) != 3 {
		err("base: 3 arguments required")
		return
	}

	sourceBase, e := strconv.Atoi(items[0])
	if e != nil {
		err("base: sourceBase not a valid integer")
		return
	}

	destBase, e := strconv.Atoi(items[1])
	if e != nil {
		err("base: destBase not a valid integer")
		return
	}

	value := items[2]

	if sourceBase == destBase {
		fmt.Println(value)
	} else if destBase == 10 {
		toBase10(value, sourceBase)
	} else if sourceBase == 10 {
		v, e := strconv.Atoi(value)
		if e != nil {
			err("base: value not a valid base10 value")
			return
		}
		fromBase10(int(v), destBase)
	} else {
		fromBase10(toBase10(value, sourceBase), destBase)
	}
}

// charToValue returns the numeric value of a digit rune
func charToValue(in rune) int {
	c := int(byte(unicode.ToUpper(in)))
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}
	return int(c - 'A' + 10)
}

// valueToChar returns the digit rune of a numeric value
func valueToChar(in int) rune {
	if in >= 0 && in <= 9 {
		return rune(in + '0')
	}
	return rune(in + 'A' - 10)
}

// fromBase10 converts a decimal number to an arbitrary base with explanation
func fromBase10(in, base int) string {
	reminders := []rune{}
	for n := in; n > 0; n /= base {
		r := n % base
		fmt.Printf("%d/%d=%d A%d\n", n, base, n/base, r)
		reminders = append([]rune{valueToChar(r)}, reminders...)
	}

	out := string(reminders)
	fmt.Println(out)
	return out
}

// toBase10 converts a number of an arbitrary base to decimal with explanation
func toBase10(in string, base int) int {
	value := 0
	parts := []string{}

	for i, c := range in {
		v := charToValue(c)
		position := len(in) - 1 - i

		value = value*base + v
		parts = append(parts, fmt.Sprintf("%d*%d^%d", v, base, position))
	}

	fmt.Printf("%s = %d\n", strings.Join(parts, " + "), value)
	return value
}

// cmd2c calculates two's complement for a binary number
func cmd2c(items []string) {
	// Items: value, bytes
	if len(items) != 2 {
		err("2c: 2 arguments required")
		return
	}

	value := items[0]

	bytes, e := strconv.Atoi(items[1])
	if e != nil {
		err("2c: bytes not a valid integer")
		return
	}

	padded := padWith0s(value, bytes)
	fmt.Print(padded + "\n")

	flipped := invBinStr(padded)
	fmt.Print(flipped + "\n\n")

	fmt.Println(binAdd(flipped, "1", bytes))
}

// padWith0s pads a string with zeros to a specific amount of bytes
func padWith0s(in string, bytes int) string {
	paddedLen := bytes * 8
	out := strings.Repeat("0", paddedLen-len(in)) + in
	return out
}

// binAdd calculates binary addition
func binAdd(a string, b string, bytes int) string {
	aPadded := padWith0s(a, bytes)
	bPadded := padWith0s(b, bytes)

	width := bytes * 8
	carry := 0
	out := make([]byte, width)

	for i := width - 1; i >= 0; i-- {
		sum := int(aPadded[i]-'0') + int(bPadded[i]-'0') + carry
		out[i] = byte('0' + sum%2)
		carry = sum / 2
	}

	fmt.Println(aPadded)
	fmt.Println(bPadded)
	fmt.Println(strings.Repeat("-", width))
	return string(out)
}

// invBinStr flips 1s and 0s in a string
func invBinStr(in string) string {
	out := []rune{}
	for _, c := range in {
		switch c {
		case '0':
			out = append(out, '1')
		case '1':
			out = append(out, '0')
		default:
			out = append(out, c)
		}
	}
	return string(out)
}

// cmd2cd calculates two's complement for a decimal number
func cmd2cd(items []string) {
	// Items: value, bytes
	if len(items) != 2 {
		err("2cd: 2 arguments required")
		return
	}

	value, e := strconv.Atoi(items[0])
	if e != nil {
		err("2cd: value not a valid integer")
		return
	}

	bytes, e := strconv.Atoi(items[1])
	if e != nil {
		err("2cd: bytes not a valid integer")
		return
	}

	if bytes <= 0 {
		err("2cd: bytes must be positive")
		return
	}

	if value < 0 {
		bits := bytes * 8
		fmt.Printf("2^%d%d=%d\n", bits, value, (1<<bits)+value)
	} else {
		fmt.Println(value)
	}
}
