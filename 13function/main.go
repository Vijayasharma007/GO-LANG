package main
import ("fmt")

/* func FunctionName() {
	// code to be executed       ---- syntax
	} */

/* func myMessage() {
	fmt.Println("I just got executed!")
  }
  
  func main() {
	myMessage() // call the function
  } */


/* func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
} */

/* func familyName(fname string){
	fmt.Println("Hello", fname)
}
func main(){
	familyName("liam")
	familyName("jenny")
} */

func familyName(fname string, age int){
	fmt.Println("hello", age, "year old", fname)
}

func main(){
	familyName("liam", 3)
	familyName("jenny", 14)
}




