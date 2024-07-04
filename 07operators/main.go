package main
import ("fmt")

func main(){

	// sum operators 
	a:= 10
	b:= 15
	c:= 10.10
	sum := a+b
	fmt.Println(sum)

	firstName := "xxx"
	lastName := "yyy"
	fmt.Printf(" %s %s \n ", firstName, lastName)

	//increment and decrement operator
	a ++
	fmt.Println("increment:", a)
	a --
	fmt.Println("decrement:", a)

	//Difference operators
	difference := a - b
	fmt.Println("Difference:", difference)

	//Product operator
	product := a*b
	fmt.Println("Product:", product)

	//Division operator
	Division := a/b
	fmt.Println("Division:", Division)

	//Remainder operator
	remainder := a%b
	fmt.Println("remainder:", remainder)

	//Type convension
	newvar := a + int(c)
	fmt.Println("newvar:", newvar)
}