package main

func romanToInt(s string) int {
    symbolToValue := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    result := 0
    n := len(s)

    for i := 0; i < n; i++ {
        currentValue := symbolToValue[s[i]]
        if i+1 < n && currentValue < symbolToValue[s[i+1]] {
            result -= currentValue
        } else {
            result += currentValue
        }
    }

    return result
}

