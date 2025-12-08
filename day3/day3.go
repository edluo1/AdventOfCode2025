package main

import ( 
    "fmt" 
    "bufio" 
    "os" 
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func getJoltages(f string) ([][]int) {
    file, err := os.Open(f)
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    var allJoltages [][]int
    for scanner.Scan() {
        batteries := scanner.Text()
        var joltages []int
        for _, c := range(batteries) {
            v := int(c - '0')
            joltages = append(joltages, v)
        }
        allJoltages = append(allJoltages, joltages)
    }

    return allJoltages
}

/*
    P1: 
    Search for the highest number and its index.
    Then look for the highest number to the right.
*/

// Input: Joltage ratings in a row.
// Output: sum of the maximum joltage ratings 
// Find the largest number you can form with two batteries, keeping them in order.
func part1(n1 [][]int) int {
    sum := 0

    for _, joltages := range n1 {
        minIdx := 0
        maxJoltage := 0
        // Search all EXCEPT the last joltage
        for idx, j := range joltages[:len(joltages)-1] {
            if j > maxJoltage {
                maxJoltage = j
                minIdx = idx
            }
        }
        // Now search for the next highest joltage
        secondMax := 0
        for idx := minIdx + 1; idx < len(joltages); idx++ {
            if joltages[idx] > secondMax {
                secondMax = joltages[idx]
            }
        }
        j := 10 * maxJoltage + secondMax
        // fmt.Printf("joltage: %v\n", j)
        sum += j
    }

    return sum
}

// Input: Joltage ratings in a row.
// Output: sum of the maximum joltage ratings 
// Find the largest number you can form with TWELVE batteries, keeping them in order.
func part2(n1 [][]int) int64 {
    var sum int64
    sum = 0
    minIdx := 0
    batteryCount := 12

    for _, joltages := range n1 {
        minIdx = 0
        j := int64(0) // The joltage.
        for b := batteryCount; b > 0; b-- {
            j *= 10

            maxJoltage := 0
            // Search from minIdx to the furthest we can go while still having space
            for idx := minIdx; idx < len(joltages)-(b-1); idx++ {
                if joltages[idx] > maxJoltage {
                    maxJoltage = joltages[idx]
                    minIdx = idx
                }
            }

            minIdx++
            j += int64(maxJoltage)
        }
        // fmt.Printf("joltage: %v\n", j)
        sum += j
    }

    return sum
 }

func main() {
    re := getJoltages("example.txt")
    rs := getJoltages("input.txt")

    fmt.Printf("example part1: %v\n", part1(re)) // example expects 357
    fmt.Printf("example part2: %v\n", part2(re)) // example expects 3121910778619
    fmt.Printf("part1: %v\n", part1(rs))
    fmt.Printf("part2: %v\n", part2(rs))
}
