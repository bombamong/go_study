# Thread

한번에 한번의 명령 수행 가능한, 하나의 명령 -> 쓰레드

멀티쓰레딩 -> 여러개의 cpu 코어에다가 쓰레드 하나씩 물려주기

CPU 코어가 하나면 ? 여러 명령어를 왔다갔다 하면서 수행 해야됨

멀리코어여도 여러 명령어를 왔다갔다 하면서 수행 함 

OS 가 쓰레드를 관리해줌



**Context Switching** -> 쓰레드 변경

쓰레드가 많으면 context switching 이 자주 발생해서 비효율적임



### GO THREAD

OS thread -> Kernel 을 한번 포장해서 만든게 Go thread

context switching 을 최소화 시킴

Go thread 는 OS thread 에 1:1 대응하지 않음



프로그램 -> 프로세스 -> 쓰레드 

#### MULTI THREADING 문제

- 동기화 문제 - Synchronizing

  해결 방법 중 하나 : mutex

  - sync.Mutex{} Lock and Unlock

- Deadlock 

  - lock 으로 인해 프로그램 진행이 안되는 상황
  - 이유: 
    - 쓰레드끼리 서로 락 잡힌걸 놓아주길 대기하는 상태

#### CHANNEL

* thread 간의 작업을 담고 있는 구조체 (queue)

* Thread safe, fixed size

  ```go
  var a chan int
  a := make(chan int, size) // fixed size
  c <- 10 push
  v := <- c
  ```



#### Producer - consumer (converyor belt) 패턴

```go
package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}

func MakeTire(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.val += "Tire, "
		outChan <- car
	}
}

func MakeEngine(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.val += "Engine"
		outChan <- car
	}
}

func StartWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car[" + strconv.Itoa(i) + "]: "}
		i++
	}
}

func main() {
	chan1 := make(chan Car)
	chan2 := make(chan Car)
	chan3 := make(chan Car)

	go StartWork(chan1)
	go MakeTire(chan1, chan2)
	go MakeEngine(chan2, chan3)
	for {
		result := <-chan3
		fmt.Println(result)
	}
}
```



## 왜 multithreading 을 해야되는가

최대의 효율을 끌어내기 위해



### Select

Switch Case 와 비슷함

```go
package main

import (
	"fmt"
	"strconv"
	"time"
)

type Car struct {
	val string
}
type Plane struct {
	val string
}

func MakeTire(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Tire_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Tire_P, "
			outPlaneChan <- plane
		}
	}
}

func MakeEngine(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Engine_P, "
			outPlaneChan <- plane
		}
	}
}

func StartCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car[" + strconv.Itoa(i) + "]: "}
		i++
	}
}

func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane[" + strconv.Itoa(i) + "]: "}
		i++
	}
}

func main() {
	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)

	go StartCarWork(carChan1)
	go StartPlaneWork(planeChan1)
	go MakeTire(carChan1, planeChan1, carChan2, planeChan2)
	go MakeEngine(carChan2, planeChan2, carChan3, planeChan3)

	for {
		select {
		case result := <-carChan3:
			fmt.Println(result)
		case result := <-planeChan3:
			fmt.Println(result)
		}
	}

}
```



