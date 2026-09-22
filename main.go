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

func err(text string) {
    fmt.Print("error: " + text + "\n")
}

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

func cmdBase(items []string) {
    // Items: sourceBase, destBase, value
    if (len(items) != 3) {
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

    if (sourceBase == 2 && destBase == 10) {
        cmdBd(value)
    }
    if (sourceBase == 10 && destBase == 2) {
        cmdDb(value)
    }
}

func cmdBd(in string) {
    value := 0
    parts := []string{}

    for i, c := range in {
        if c != '0' && c != '1' {
            err("bd: invalid binary number")
            return
        }

        bit := int(c - '0')
        position := len(in) - 1 - i
        power := 1 << position

        value += bit * power
        parts = append(parts, fmt.Sprintf("%d*2^%d", bit, position))
    }

    fmt.Printf("%s = %d\n", strings.Join(parts, " + "), value)
}

func cmdDb(in string) {
    value, e := strconv.Atoi(in)
    if e != nil {
        err("2cd: value not a valid integer")
        value = 0
    }

    reminders := []int{}

    for n := value; n > 0; n /= 2 {
        remainder := n % 2
        fmt.Printf("%d/2=%d A%d\n", n, n / 2, remainder)
        reminders = append(reminders, remainder)
    }

    for i := len(reminders) - 1; i >= 0; i-- {
        fmt.Print(reminders[i])
    }
    fmt.Println()
}

func charValue(l rune) int {
    c := int(byte(unicode.ToUpper(l)))
    if c >= '0' && c <= '9' {
        return int(c - 48)
    }
    return int(c - 64);
}

func cmd2c(items []string) {
    // Items: value, bytes
    if (len(items) != 2) {
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

func padWith0s(in string, bytes int) string {
    paddedLen := bytes * 8
    out := strings.Repeat("0", paddedLen - len(in)) + in
    return out
}

func binAdd(a string, b string, bytes int) string {
    aPadded := padWith0s(a, bytes)
    bPadded := padWith0s(b, bytes)

    width := bytes * 8
    carry := 0
    out := make([]byte, width)

    for i := width - 1; i >= 0; i-- {
        sum := int(aPadded[i] - '0') + int(bPadded[i] - '0') + carry
        out[i] = byte('0' + sum % 2)
        carry = sum / 2
    }

    fmt.Println(aPadded)
    fmt.Println(bPadded)
    fmt.Println(strings.Repeat("-", width))
    return string(out)
}

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

func cmd2cd(items []string) {
    // Items: value, bytes
    if (len(items) != 2) {
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
