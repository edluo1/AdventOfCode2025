package main

import ( 
    "fmt" 
    "bufio" 
    "os" 
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

const LEFT byte = 'L'
const RIGHT byte = 'R'

type rotation struct {
    direction byte
    turn int
}

func getRotations(f string) ([]rotation) {
    file, err := os.Open(f)
    check(err)
    defer file.Close()

    scanner:= bufio.NewScanner(file)

    var rs []rotation
    for scanner.Scan() {
        t := scanner.Text()
        d := t[0]
        n1, err := strconv.Atoi(t[1:])
        check(err)

        r := rotation{d, n1}

        rs = append(rs, r)
    }

    return rs
}

func mod(x, y int) int {
    m := x % y

    if (m < 0) {
        m += y
    }

    return m
}

// Input: Set of rotations. L subtracts from dial, R adds to dial.
// Output: Number of times the dial points to 0 after stopping.
func part1(n1 []rotation) int {
    dial := 50
    zeroCount := 0

    for _, r := range n1 {
        if r.direction == LEFT {
            dial -= r.turn
        } else {
            dial += r.turn
        }
        dial = mod(dial, 100)

        if dial == 0 {
            zeroCount += 1
        }
        // fmt.Println(dial)
    }

    return zeroCount
}

// Input: Set of rotations. L subtracts from dial, R adds to dial.
// Output: Number of times the dial points to 0 at any time during a rotation.
func part2(n1 []rotation) int {
    dial := 50
    zeroCount := 0

    for _, r := range n1 {
        startPos := dial
        endPos := dial
        if r.direction == LEFT {
            endPos -= r.turn
        } else {
            endPos += r.turn
        }
        dial = mod(endPos, 100)

        if (endPos <= 0) { // left turn
            // 5 -> -100 = 2 0s
            // 5 -> -5 = 1 0
            // 0 -> -105 = 1 0
            zeroCount += (-endPos / 100) + 1
            if (startPos == 0) {
                zeroCount -= 1
            }
        } else if (endPos >= 100) {
            // 5 -> 100 = 1 0
            // 0 -> 210 = 2 0s
            // 91 -> 300 = 3 0s
            zeroCount += endPos / 100
        }

        // fmt.Println(dial)
    }

    return zeroCount
}

func main() {
    re := getRotations("example.txt")
    rs := getRotations("input.txt")

    fmt.Printf("part1: %v\n", part1(re)) // example expects 3
    fmt.Printf("part2: %v\n", part2(re)) // example expects 6
    fmt.Printf("part1: %v\n", part1(rs))
    fmt.Printf("part2: %v\n", part2(rs))
}

