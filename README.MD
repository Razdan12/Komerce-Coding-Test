# Komerce Coding Test

## Problem 1 : Sort Character

### Input :

One line of words (S).

### Output :

Contains vowel and consonant characters that has been sorted according to the following rules.

- Sort the letters according to the order they came out.
- Separate between vowels and consonants.
- Process as lowercase letters (whitespaces are ignored).

### Test case :

### Example Input 1

```bash
Input one line of words (S) : Sample Case
```

### Example Output 1

```bash
Vowel Characters : aaee
Consonant Characters : ssmplc
```

### Example Input 2

```bash
Input one line of words (S) : Next Case
```

### Example Output 2

```bash
Vowel Characters : eea
Consonant Characters : nxtcs
```

## Problem 1 Answer

### With Golang

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func sortCharacters(s string) (string, string) {
	vowels := "aiueo"
	var vowelsCars, consonanChart string

	s = strings.ToLower(s)
    for _, char := range s {
        if strings.ContainsRune(vowels, char){
            vowelsCars += string(char)
        }else if char >= 'a' && char <= 'z' {
            consonanChart += string(char)
        }
    }
    return vowelsCars, consonanChart
}
func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a string: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    vowels, consonants := sortCharacters(input)
    fmt.Println("Vowel Characters:", vowels)
    fmt.Println("Consonant Characters:", consonants)
}
```

### With Javascript

```js
const sortCharacters = (s) => {
  const vowels = "aeiou";
  let vowelChars = "";
  let consonantChars = "";
  s = s.toLowerCase();
  for (let char of s) {
    if (vowels.includes(char)) {
      vowelChars += char;
    } else if (char >= "a" && char <= "z") {
      consonantChars += char;
    }
  }
  return { vowelChars, consonantChars };
};

const input = "Sample Case";
const { vowelChars, consonantChars } = sortCharacters(input);
console.log("Vowel Characters:", vowelChars);
console.log("Consonant Characters:", consonantChars);
```

## Problem 2 : PSBB ( Pembatasan Sosial Berskala Besar )

Galih and Ratna married during the COVID 19 period and only invited the families of both
partners. they rented a number of minibuses to pick up all of their families to go to the wedding.
But during COVID 19, the government held a PSBB program to reduce the impact of the spread
of the virus. Each mini bus can only carry at most 4 passengers.
What a minimum number of buses will they need to rent if all members of each family should ride
in the same Busses. (one bus can't take more than two family)

### Input

The first line contains integer n — the number of families.
The second line contains a sequence of integers. The integers are separated by a space, each
integer is the number of members in the family.

### Output

Print the single number — the minimum number of buses necessary to drive all family to the
Wedding.
Print “ Input must be equal with count of family” if input number of family is not equal with count
of family.

### Example Input 1

```bash
Input the number of families : 5
Input the number of members in the
family ( separated by a space) : 1 2 4 3 3
```

### Example Output 1

```bash
Minimum bus required is : 4
```

### Example Input 2

```bash
Input the number of families : 8
Input the number of members in the
family ( separated by a space) : 2 3 4 4 2 1 3 1
```

### Example Output 2

```bash
Minimum bus required is : 5
```

### Example Input 3

```bash
Input the number of families : 5
Input the number of members in the
family ( separated by a space) : 1 5
```

### Example Output 3

```bash
Input must be equal with count of family
```

## Problem 2 Answer

### With Golang

```go
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


```

### With Javascript

```js
const minBuses = (n, members) => {
  if (members.length !== n) {
    console.log("Input must be equal with count of family");
    return;
  }

  members.sort((a, b) => b - a);
  let buses = 0;
  let i = 0,
    j = members.length - 1;

  while (i <= j) {
    if (members[i] + members[j] <= 4) {
      j--;
    }
    i++;
    buses++;
  }

  console.log("Minimum bus required is:", buses);
};

const n = 5;
const members = [1, 2, 4, 3, 3];
minBuses(n, members);
```
