# 체육복

> 문제 설명
> 점심시간에 도둑이 들어, 일부 학생이 체육복을 도난당했습니다. 다행히 여벌 체육복이 있는 학생이 이들에게 체육복을 빌려주려 합니다. 학생들의 번호는 체격 순으로 매겨져 있어, 바로 앞번호의 학생이나 바로 뒷번호의 학생에게만 체육복을 빌려줄 수 있습니다. 예를 들어, 4번 학생은 3번 학생이나 5번 학생에게만 체육복을 빌려줄 수 있습니다. 체육복이 없으면 수업을 들을 수 없기 때문에 체육복을 적절히 빌려 최대한 많은 학생이 체육수업을 들어야 합니다.

> 전체 학생의 수 n, 체육복을 도난당한 학생들의 번호가 담긴 배열 lost, 여벌의 체육복을 가져온 학생들의 번호가 담긴 배열 reserve가 매개변수로 주어질 때, 체육수업을 들을 수 있는 학생의 최댓값을 return 하도록 solution 함수를 작성해주세요.

> 제한사항
> 전체 학생의 수는 2명 이상 30명 이하입니다.
> 체육복을 도난당한 학생의 수는 1명 이상 n명 이하이고 중복되는 번호는 없습니다.
> 여벌의 체육복을 가져온 학생의 수는 1명 이상 n명 이하이고 중복되는 번호는 없습니다.
> 여벌 체육복이 있는 학생만 다른 학생에게 체육복을 빌려줄 수 있습니다.
> 여벌 체육복을 가져온 학생이 체육복을 도난당했을 수 있습니다. 이때 이 학생은 체육복을 하나만 도난당했다고 가정하며, 남은 체육복이 하나이기에 다른 학생에게는 체육복을 빌려줄 수 없습니다.

## My solution

**_GO version_**

```go
func solution(n int, lost []int, reserve []int) int {
    // 참여 못하는 학생 수를 담는 변수
    out := len(lost)

    // 여벌이 있지만, 동시에 잃어버린 학생들은 열외함과 동시에
    // 참여 학생 1 증가
    for i, r := range reserve {
        idx := indexOf(lost, r)
        if idx != -1 {
            out--
            lost[idx] = -1
            reserve[i] = -1
        }
    }

    // lost 와 reserve 의 값이 1이 다르면 빌려주고
    // 참여 학생 1 증가 하고
    // 또 받을 수 없게 처리
    for _, r := range reserve {
        for li, l := range lost {
            if r - l == 1 || r - l == -1 {
                out--
                lost[li] = -1
                break
            }
        }

    }

    return n - out
}

func indexOf(s []int, d int) int {
    for i, v := range s {
        if v == d {
            return i
        }
    }
    return  -1
}
```

**_JS version_**

```js
function solution(n, lost, reserve) {
  // 배열에 잃어버린 사람은 0, 여벌 가지고 있는 사람은 2로 변경
  let answer = new Array(n + 1).fill(1);
  lost.forEach(e => answer[e]--);
  reserve.forEach(e => answer[e]++);

  // 배열 0번째 제거, 학생은 1부터 시작함
  answer[0] = 0;

  // 체육복이 없고 옆에사람 두개 있으면 뺏어오기
  for (let idx = 1; idx < answer.length; idx++) {
    if (answer[idx] < 1) {
      if (answer[idx - 1] > 1) {
        answer[idx - 1]--;
        answer[idx]++;
      } else if (answer[idx + 1] > 1) {
        answer[idx + 1]--;
        answer[idx]++;
      }
    }
  }

  // 체육복 있는 사람 다 더하기, 지금까지 두개 가지고 있는 사람은 못 나눠줬기 때문에 1인분
  answer = answer.reduce((acc, curr) => (acc += curr > 1 ? 1 : curr));
  return answer;
}
```
