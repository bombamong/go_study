## Slice

C++ vector

Java ArrayList

C# List

Pythong Slice

Golang Slice

length vs cap

들어있는 내용 길이 

실제 배열 길이

make([]int, 2 length, 3 cap)

append(mySlice, 3) --> cap 없으면 새걸로 리턴



### Slice a slice

mySlice[startIndex: endIndex] not inclusive of endIndex

mySlice[3 : ] omission of endIndex --> to end

mySlice[: 3] omission of startIndex --> from start



