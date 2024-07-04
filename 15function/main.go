package main
import ("fmt")

func main() {
	fmt.Println("welcome to function in golang")
	greeter()
	greeterTwo()
	result := adder(3,4)
	fmt.Println("Result is:", result)

	
	
}

func adder(valOne int, valTwo int)int{
	return valOne + valTwo
}

func greeter(){
	fmt.Println("Hello from golang")
}

func greeterTwo(){
	fmt.Println("This is another method")
}