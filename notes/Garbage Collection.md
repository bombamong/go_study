## Garbage Collection

메모리를 지워줘야됨

C에서는...

malloc(size) returns address (pointer)

Memory Leak --> 메모리 잡아놓고 free 못하게 만들어 둔거

그래서 나온게 Garbage Collector

refrence count 가 0 이 되면 지움



Dangling --> 없어진 변수 주소를 참조하고 있는거

C 에서는 스택 vs 힙 메모리 구분 필요함

Go 에서는 밖에서 봤을때는 구분 없음



외딴섬

a --> b --> c --> a

지들끼리 가르키고 있는데 다른 아무것도 얘네들을 찾아주지 않음

memory leak.

