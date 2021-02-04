package main

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
	(a).color = c
}
