package main
import ("fmt")

func main(){
	age := 21
	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 13 {
		fmt.Println("You are teenager")
	}else {
		fmt.Println("You are child")
	}
} 


