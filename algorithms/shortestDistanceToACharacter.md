# Shortest Distance to a Character

**_PROBLEM FROM LEETCODE_**

> Given a string `s` and a character `c` that occurs in `s`, return *an array of integers `answer` where* `answer.length == s.length` *and* `answer[i]` *is the shortest distance from* `s[i]` *to the character* `c` *in* `s`. 
>
> **Example 1:**
>
> ```
> Input: s = "loveleetcode", c = "e"
> Output: [3,2,1,0,1,0,0,1,2,2,1,0]
> ```
>
> **Example 2:**
>
> ```
> Input: s = "aaab", c = "b"
> Output: [3,2,1,0]
> ```
>
> **Constraints:**
>
> - `1 <= s.length <= 104`
> - ```s[i]` and `c` are lowercase English letters.
> - `c` occurs at least once in `s`.

### My Solution

```go
func shortestToChar(s string, c byte) []int {
    answer := make([]int, len(s))
    sSlice := []byte(s)
    
    for i, v := range sSlice {
        if v == c {
            answer[i] = 0
        } else {
          	// Search left side of slice and get the min distance to c
            left := searchLeft(i, len(sSlice), sSlice, c)
          	// Search right side of slice and get the min distance to c
            right := searchRight(i, len(sSlice), sSlice, c)            
          	// assign the shortest distance to answer[i]
            if left < right {
                answer[i] = left
            } else {
                answer[i] = right  
            }
        }
    }
    
    return answer
}

func searchLeft(i int, maxDistance int, searchSlice []byte, c byte) int {
    distance := maxDistance
    currentIdx := i
    for currentIdx >= 0 {
        if searchSlice[currentIdx] == c {
            if currentIdx > i {
                distance = currentIdx - i
            } else {
                distance = i - currentIdx
            }
         break;
        }
    currentIdx--
    }
    return distance
}

func searchRight(i int, maxDistance int, searchSlice []byte, c byte) int {
    distance := maxDistance
    currentIdx := i
    for currentIdx < maxDistance {
        if searchSlice[currentIdx] == c {
            if currentIdx > i {
                distance = currentIdx - i
            } else {
                distance = i - currentIdx
            }
         break;
        }
    currentIdx++
    }
    return distance
}
```