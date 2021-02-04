## Notes for Golang

1. ; 안씀

2. ( ) 안씀

3. while 없음 -> for 만 사용

   JS:

   ```javascript
   while (something) {
     doSomething();
   }

   //OR

   do {
     doSomething();
   } while (something);
   ```

   Go:

   ```go
   for {
     if(something) break;
     doSomething()
   }

   //OR

   for {
     doSomething()
     if(something) break;
   }
   ```

4. func -> can have multiple returns

   Single return

   Java:

   ```java
   int addTwoIntegers(int num1, int num2) {
     return (num1 + num2)
   }
   ```

   Go:

   ```go
   func addTwoIntegers(num1, num2 int) int {
     return (num1 + num2)
   }
   ```

   Multiple returns

   Go:

   ```go
   func returnIntAndFloat(num1 int, num2 float64) (int, float64) {
   	return num1, num2
   }

   a, b := returnIntAndFloat(1, 1.1)
   fmt.Println(a, b) // 1 1.1
   ```

5. 문자열

   Go 는 기본으로 UTF-8 사용하고 있음

   ASCII 처럼 1 BYTE 가 아니기 떄문에 눈으로 보이는것과 길이가 다를수 있음 (len(s) 사용 시)

   문자열 다룰때 []byte 로 다룰수도 있고 []rune 으로 다룰수도 있다.

   byte 나는 마그대로 8bit 마다 끊기고 rune 은 1~3 바이트 중 하나의 character 를 나타내는곳마다 잘린다.

   ```go
   s := "Hello 안녕"

   //[]byte
   for i := 0; i < len(s); i++ {
     fmt.Print(string(s[i]), ", ")
   }
   fmt.Println()
   ```


   //[]rune
   for _, c := range s {
     fmt.Print(string(c), ", ")
   }

   // or convert to []rune
   s2 := []rune(s)
   for i := 0; i < len(s2); i++ {
     fmt.Print(string(s2[i]), ", ")
   }
   fmt.Println()


   ```

6. Array is immutable

   Javascript

   ```js
   0let arr = [1,2,3,4,5]
   let arr2 = arr
   arr2[0] = 2
   console.log(arr, arr2) // [2, 2, 3, 4, 5] [2, 2, 3, 4, 5]
   ```

   Go

   ```go
   	arr := [5]int{1, 2, 3, 4, 5}
   	arr2 := arr
   	arr2[0] = 2
   	fmt.Println(arr, arr2) // [1 2 3 4 5] [2 2 3 4 5]
   ```

   BUUUUUT: slices are mutable \*

   ```go
   	slice := []int{1, 2, 3, 4, 5}
   	slice2 := slice
   	slice2[0] = 2
   	fmt.Println(slice, slice2) // [2 2 3 4 5] [2 2 3 4 5]
   ```

7. Structs

   A collection of data that relates to a certain object

   ```go
   type Student struct {
     name string;
   }
   ```

   행동 추가하는 방법

   ```go
   func (s Student) printName() {
     fmt.Printf("My name is %s", s.name)
   }
   ```

   ### 중요한 부분!!

   아래와 같이 method 안에서 struct 의 데이터를 변경하는건 의미가 없다, pointer 사용해야함

   ```go
   /*
   Animal struct
   */
   type Animal struct {
   	color string
   }

   /*
   ChangeColor pretends to change Animal color
   */
   func (a Animal) ChangeColor(c string) {
   	a.color = c
   }

   /*
   ReallyChangeColor really changes Animal color
   */
   func (a *Animal) ReallyChangeColor(c string) {
   	(*a).color = c
     // this can be written as below, dereferencing done by Go
     // a.color = c
   }
   ```

   func main() {
     animal1 := Animal{color: "brown"}
     animal1.ChangeColor("blue")
     fmt.Println(animal1.color) // "brown"
   }

## POINTER

변수는 value 를 대표하고 있고,

포인터는 value 가 있는 주소를 가르키고 있다

```go
var a int
var p *int
p = &a
a = 3
fmt.Println(p) // 주소
fmt.Println(*p) // 주소에 있는 값
```