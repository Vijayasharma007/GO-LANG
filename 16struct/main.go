package main

import (
	"fmt"
)

/*While arrays are used to store multiple values of the same data type into a single variable,
structs are used to store multiple values of different data types into a single variable. */

/* type struct_name struct {
  	   member1 datatype;
 	   member2 datatype;
 	   member3 datatype;
  ...
} */

/* func main(){
type student struct{
	firstname string
	lastname string
	age int
	subject [] string
}
var student1 student
student1 = student{"go \n"  ,"language \n ",2024 ,[]string{"Tamil","english"}}
fmt.Printf("%+v", student1)
} */

func main(){

	golanguage := User{"golanguage", "golanguage@gmail.com", 2024}
	fmt.Println(golanguage)
}

type User struct{
	Name string
	Email string
	Age int
}