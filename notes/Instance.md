# Instance

```go
type Student struct {
  name string
  age int
  grade int
}

func main() {
  a := Student{"aaa", 20, 10}
  b := a
  b.age = 30
}
```

b 는 a 의 copy 임

value type vs reference type 잘 구분 하도록!

```go
type Student struct {
  name string
  age int
  grade int
}

func main() {
  a := Student{"aaa", 20, 10}
  b := &a
  b.age = 30
}
```

b 는 a 의 주소를 가르키고 있음 (*b).age 이게 맞지만 go에서는 알아서 해줌

