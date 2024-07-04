package main
import ("fmt")

func main(){

	//you can declare the variable and assign a value to it on one line:
	
  var student1 string = "John" //type is string
  var student2 = "Jane" //type is inferred
  x := 2 //type is inferred

  fmt.Println(student1)
  fmt.Println(student2)
  fmt.Println(x)
} 

/* func main(){
	//if you declare a variable without an initial value, its value will be set to the default value of its type
  var a string
  var b int
  var c bool

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)

} */

/* func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
} */



