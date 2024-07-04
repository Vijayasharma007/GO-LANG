package main
import ("fmt")

/* func main() {

	// Use the if statement to specify a block of Go code to be executed if a condition is true.
	/* if condition {
		// code to be executed if condition is true
	  } */

	/* if 10 > 20 {
		fmt.Println("30 is greater than 20")
	} */



	// Use the else statement to specify a block of code to be executed if the condition is false
	/* syntax

	 if condition {
		// code to be executed if condition is true
	  } else {
		// code to be executed if condition is false
	  }  */

/* func main(){
a := 10
if (a > 20) {
	fmt.Println("this if condition is true:", a)
}else{
	fmt.Println("this if condition is false:", a)
}
} */



// Use the else if statement to specify a new condition if the first condition is false

/* if condition1 {
	// code to be executed if condition1 is true
 } else if condition2 {
	// code to be executed if condition1 is false and condition2 is true
 } else {
	// code to be executed if condition1 and condition2 are both false
 }  */


 /* func main() {
	time := 22
	if time < 10 {
	  fmt.Println("Good morning.")
	} else if time < 20 {
	  fmt.Println("Good day.")
	} else {
	  fmt.Println("Good evening.")
	}
  } */

// You can have if statements inside if statements, this is called a nested if.

/* if condition1 {
	// code to be executed if condition1 is true
   if condition2 {
	  // code to be executed if both condition1 and condition2 are true
   }
 }  */


func main() {
   num := 20
   if num >= 10 {
	 fmt.Println("Num is more than 10.")
	 if num > 15 {
	   fmt.Println("Num is also more than 15.")
	  }
   } else {
	 fmt.Println("Num is less than 10.")
   }
 }

	 
	