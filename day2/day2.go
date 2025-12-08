package main

import ( 
    "fmt" 
    "bufio" 
    "os" 
    "math"
    "strconv"
    s "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}


// NOTE: the ranges are inclusive.
type id_range struct {
    start int64
    end int64
}

func getIDRanges(f string) ([]id_range) {
    file, err := os.Open(f)
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var ranges []id_range
    for scanner.Scan() {
        split := s.Split(scanner.Text(), ",")
        for _, elem := range(split) {
            vals := s.Split(elem, "-")
            val1, err := strconv.ParseInt(vals[0], 10, 64)
            check(err)
            val2, err := strconv.ParseInt(vals[1], 10, 64)
            check(err)
            idr := id_range{val1, val2}
            ranges = append(ranges, idr)
        }
    }

    return ranges
}

func getDigits(n int64) int64 {
    return int64(math.Log10(float64(n))) + 1
}

/*
    P1: 
    Only IDs with even digits can be invalid.
    2-digit ones are divisible by 11, 4-digit ones are divisible by 101, etc.
*/

// Input: List of ID ranges that you need to check. The ranges are inclusive.
// Output: Sum of all invalid IDs; that is: IDs that are just one number repeated twice.
func part1(n1 []id_range) int64 {
    sum := int64(0)

    for _, r := range n1 {
        for i := r.start; i <= r.end; i++ {
            digits := getDigits(i)
            if digits % 2 == 0 {
                multiple := int64(math.Pow(float64(10), float64(digits / 2))) + 1
                if i % multiple == 0 {
                    sum += i
                }
            }
        }
    }

    return sum
}

// digitCount: 1 -> 1, 2 -> 10, 3 -> 100
func createMultiple(digitCount int64, repeats int64) int64 {
    multiple := int64(0)
    for i := int64(0); i < repeats; i++ {
        multiple *= int64(math.Pow(float64(10), float64(digitCount)))
        multiple += 1
    }
    return multiple
}

/* A 9-digit number repeats 3 times and is a multiple of 1001001.
 1001001
*/

// Input: List of ID ranges that you need to check. The ranges are inclusive.
// Output: Sum of all invalid IDs; that is: IDs that are just one number repeated 2 or MORE times.
func part2(n1 []id_range) int64 {
    sum := int64(0)

    for _, r := range n1 {
        for i := r.start; i <= r.end; i++ {
            digits := getDigits(i)
            for r := int64(2); r <= digits; r++ {
                if digits % r == 0 {
                    repeatCount := digits / r
                    multiple := createMultiple(repeatCount, r)

                    if i % multiple == 0 {
                        fmt.Println(i)
                        sum += i
                        break
                    }
                }
            }
        }
    }

    return sum
}

func main() {
    re := getIDRanges("example.txt")
    rs := getIDRanges("input.txt")

    fmt.Printf("example part1: %v\n", part1(re)) // example expects 1227775554
    fmt.Printf("example part2: %v\n", part2(re)) // example expects 4174379265
    fmt.Printf("part1: %v\n", part1(rs))
    fmt.Printf("part2: %v\n", part2(rs))
}
