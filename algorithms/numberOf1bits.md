**_PROBLEM FROM LEETCODE_**

> Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).
> Note:
> Note that in some languages such as Java, there is no unsigned integer type. In this case, the input will be given as a signed integer type. It should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
> In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 3 above, the input represents the signed integer. -3.
> Follow up: If this function is called many times, how would you optimize it?

**Example 1:**

```
Input: n = 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.
```

**Example 2:**

```
Input: n = 00000000000000000000000010000000
Output: 1
Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit.
```

**Example 3:**

```
Input: n = 11111111111111111111111111111101
Output: 31
Explanation: The input binary string 11111111111111111111111111111101 has a total of thirty one '1' bits.
```

**Constraints:**

- The input must be a **binary string** of length `32`

```go
import "strconv"

func hammingWeight(num uint32) int {
    n := strconv.FormatInt(int64(num), 2)
    counter := 0
    for _, v := range n {
        if v == '1' {
            counter++
        }
    }
    return counter
}
```

**좋은 방법이 아닌 것 같음**

## SOLUTION

```go
func hammingWeight(num uint32) int {
   counter := 0
   for num > 0 {
       num = num & (num - 1)
     counter++
   }
   return counter
}
```

### 설명

> 예) input 15 또는 1111 \*32bit int 지만 일단 무시

위와 같이 bitwise AND "&" 을 사용했을때 1 vs 1 아 아닌 경우 전부 0 을 리턴 하게 됨.

즉

n & (n - 1) 을 했을 경우 n 이 가진 가장 작은 1 이 0으로 변경 되게 된다.

예를 들어

15 & 14 = 14

111**1**

1110 &

111**0**

15의 가장 오른쪽 끝에 있는 1이 0으로 변경되면서 14가 나온다

반복해서

14 & 13 = 12

11**1**0

1101 &

11**0**0

12 & 11 = 8

1**1**00

1011 &

1**0**00

8 & 7 = 0

**1**000

0111 &

**0**000

위와 같이 초기 번호인 15가 0이 될때까지 그 숫자가 가진 가장 작은 1을 0으로 변환 시켜주기때문에

몇번만에 0이 됐는지를 count 에 저장해주면 1이 몇개가 있었는지 알 수 있다.
