# Interface

객체의 기능이 정의 된 것

```go
type A struct{
  status string  
}

func (a *A) MethodA() {
  A.ExternalB()
}
```

**Example**

```go
package main

import "fmt"

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type SpoonOfJam interface {
	String() string
}

type OrangeJam struct {
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

type SpoonOfOrangeJam struct {
}

func (spoon *SpoonOfOrangeJam) String() string {
	return " + one spoon of orange jam"
}

type AppleJam struct {
}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

type SpoonOfAppleJam struct {
}

func (spoon *SpoonOfAppleJam) String() string {
	return " + one spoon of apple jam"
}

type Bread struct {
	val string
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func main() {
	bread := &Bread{val: "bread"}
	jam := &OrangeJam{}
	bread.PutJam(jam)
	fmt.Println(bread.val) // bread + one spoon of orange jam
}
```

