package main
import ("fmt")

func main() {

	var a uint8 = 241
	var b uint8 = 141

	// OR operator
	c := a | b
	fmt.Println("OR operator:", c)

	//ADD operator
	d := a & b
	fmt.Println("ADD operator:", d)

	//invert operator
	e := ^a
	fmt.Println("invert:", e)
}