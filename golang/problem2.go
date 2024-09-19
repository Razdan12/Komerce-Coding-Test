package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
	"sort"
)

func minBuses(n int, members []int) int {
    if len(members) != n {
        fmt.Println("Input must be equal with count of family")
        return -1
    }

    sort.Sort(sort.Reverse(sort.IntSlice(members))) 
    buses := 0
    i, j := 0, len(members)-1

    for i <= j {
        if members[i] + members[j] <= 4 {
            j--
        }
        i++
        buses++
    }

    return buses
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Input the number of families: ")
    input, _ := reader.ReadString('\n')
    n, err := strconv.Atoi(strings.TrimSpace(input))
    if err != nil {
        fmt.Println("Invalid input")
        return
    }

    fmt.Print("Input the number of members in the family (separated by a space): ")
    input, _ = reader.ReadString('\n')
    membersStr := strings.Fields(input)
    var members []int
    for _, s := range membersStr {
        num, err := strconv.Atoi(s)
        if err != nil {
            fmt.Println("Invalid input")
            return
        }
        members = append(members, num)
    }

    result := minBuses(n, members)
    if result != -1 {
        fmt.Println("Minimum bus required is:", result)
    }
}
